package config

import (
	"os"
)

const (
	DefaultServerPort  = "8080"
	DefaultMetricsPath = "/metrics"
)

type Config struct {
	Port        string
	Template    string
	MetricsPath string
}

func New() (*Config, error) {

	cfg := Config{}
	cfg.Port = os.Getenv("SERVER_PORT")

	if cfg.Port == "" {
		cfg.Port = DefaultServerPort
	}
	cfg.MetricsPath = os.Getenv("METRICS_PATH")
	if cfg.MetricsPath == "" {
		cfg.MetricsPath = DefaultMetricsPath
	}
	cfg.Template = `<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>k8s Image Pull Exporter</title>
				<style>
					body {
						font-family: Arial, sans-serif;
						margin: 0;
						padding: 0;
						background-color: #f0f0f0;
					}
					.container {
						max-width: 800px;
						margin: 0 auto;
						padding: 20px;
					}
					h1 {
						color: #333;
						text-align: center;
						margin-bottom: 20px;
					}
					p {
						text-align: center;
						margin-bottom: 20px;
					}
					a {
						display: inline-block;
						background-color: #007bff;
						color: white;
						padding: 10px 20px;
						text-decoration: none;
						border-radius: 5px;
					}
					a:hover {
						background-color: #0056b3;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h1>k8s Image Pull Exporter</h1>
					<p>Monitor the availability of container images in your Kubernetes cluster.</p>
					<p><a href="/metrics">View Metrics</a></p>
				</div>
			</body>
			</html>`

	return &cfg, nil
}
