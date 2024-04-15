package app

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"infra/k8s-image-metrics/pkg/config"
	"infra/k8s-image-metrics/pkg/connection"
	"infra/k8s-image-metrics/pkg/metrics"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type MetricsServer struct {
	Metrics *metrics.Metrics
	Server  *http.ServeMux
	Counter int
}

func New() *MetricsServer {
	return &MetricsServer{
		Metrics: metrics.New(),
		Server:  http.NewServeMux(),
	}
}

func (ms *MetricsServer) GatherMetrics(namespaces []string, interval time.Duration) {
	for _, namespace := range namespaces {
		podList, err := GetPodsFromNamespace(namespace)
		if err != nil {
			log.Fatalf("Get pods error from namesapce, %v\n", err)
			os.Exit(1)
		}

		// Calculate the cutoff time for pods to be considered "new"
		cutoffTime := time.Now().Add(-time.Duration(interval.Minutes()) * time.Minute)
		// Filter pods based on their creation timestamp
		for _, pod := range podList.Items {
			podCreationTime := pod.GetCreationTimestamp().Time
			if podCreationTime.After(cutoffTime) {
				ms.GetPodEvents(namespace, pod.Name)
			}
		}
	}
}

func (ms *MetricsServer) Run(cfg *config.Config, namespaces []string, interval time.Duration) {
	ms.Server.Handle(cfg.MetricsPath, ms.Metrics.Handler())
	ms.Server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(cfg.Template))
	})
	go func() {
		log.Infof("Start handling server on port %v", cfg.Port)
		log.Fatal(http.ListenAndServe(":"+cfg.Port, ms.Server))
	}()
	for {
		ms.GatherMetrics(namespaces, interval)
		time.Sleep(time.Duration(interval.Seconds()) * time.Second)
	}
}

func GetPodsFromNamespace(namespace string) (*v1.PodList, error) {
	clientset, err := connection.GetK8sClient()
	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("Failed to get pods: %v", err)
	}
	return podList, nil
}

func (ms *MetricsServer) GetPodEvents(namespace, podName string) {
	clientset, err := connection.GetK8sClient()
	if err != nil {
		panic(err.Error())
	}
	events, err := clientset.CoreV1().Events(namespace).List(context.Background(), metav1.ListOptions{
		FieldSelector: fmt.Sprintf("involvedObject.name=%s,involvedObject.namespace=%s", podName, namespace),
	})
	if err != nil {
		panic(err.Error())
	}

	// Print the events
	for _, event := range events.Items {
		if event.Reason == "Pulled" {
			message := event.Message
			// Use regex to find the number representing the pull time
			rePullTime := regexp.MustCompile(`(\d+\.\d+)(ms|s)?`)
			pullTimeMatches := rePullTime.FindStringSubmatch(message)
			// Regular expression to match the base image name
			reBaseImageName := regexp.MustCompile(`[^\/]+:`)
			baseImageNameMatches := reBaseImageName.FindStringSubmatch(message)
			if len(pullTimeMatches) > 0 && len(baseImageNameMatches) > 0 {
				// Convert the matched string for pull time to a float64
				pullTime, err := strconv.ParseFloat(pullTimeMatches[1], 64)
				if err != nil {
					fmt.Println("Error parsing pull time:", err)
					return
				}
				if pullTimeMatches[2] == "ms" {
					pullTime /= 1000
				}
				pulledImageName := strings.TrimSuffix(baseImageNameMatches[0], ":")
				ms.Metrics.ImagePullDurationSecondsGauge.With(prometheus.Labels{"namespace": namespace, "pod_name": podName, "image_name": pulledImageName}).Set(pullTime)
			} else {
				fmt.Println("No pull time found in the message.")
			}
		}
	}
}
