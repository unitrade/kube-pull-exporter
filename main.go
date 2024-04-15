package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"infra/k8s-image-metrics/pkg/app"
	"infra/k8s-image-metrics/pkg/config"
	"strings"
	"time"
)

const (
	DefaultNamespace = "default"
)

var (
	namespacesFlag string
	namespaces     []string
	interval       time.Duration
)

func init() {
	flag.StringVar(&namespacesFlag, "namespaces", "", "Namespaces for image check")
	flag.DurationVar(&interval, "interval", 300*time.Second, "time interval (in seconds) to check")
	flag.Parse()
	if namespacesFlag == "" {
		namespaces = append(namespaces, DefaultNamespace)
		log.Infoln("Flag source path '-namespaces' not set. The 'default' namespace is set")
	} else {
		namespaces = strings.Split(namespacesFlag, ",")
		log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
		log.Info("Start k8s-image-metrics ...")
	}
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	server := app.New()
	server.Run(cfg, namespaces, interval)
}
