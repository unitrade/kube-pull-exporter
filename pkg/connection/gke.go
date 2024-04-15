package connection

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

var k8sClient *kubernetes.Clientset

func GetK8sClient() (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error
	if k8sClient != nil {
		return k8sClient, nil
	}

	// Check if we are running inside a Kubernetes cluster
	if _, err := os.Stat("/var/run/secrets/kubernetes.io/serviceaccount"); err == nil {
		// Running inside a Kubernetes cluster
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to create in-cluster config: %v", err)
		}
	} else {
		// Running outside a Kubernetes cluster
		home := homedir.HomeDir()
		kubeConfigPath := filepath.Join(home, ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfigPath)
		if err != nil {
			return nil, fmt.Errorf("failed to build config from a kubeconfig filepath: %v", err)
		}
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new Clientset for the given config: %v", err)
	}

	k8sClient = clientSet
	return k8sClient, nil
}
