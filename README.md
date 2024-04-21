# kube-pull-exporter
![Version: 0.0.2](https://img.shields.io/badge/Version-0.0.2-informational?style=flat-square)
![GoLang: 1.21.7](https://img.shields.io/static/v1?label=GoLang&message=1.21.7&color=purple&logo=go)
![License](https://img.shields.io/static/v1?label=License&message=Apache-2.0&color=green&logo=apache)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)
![Image size: 39.8MB](https://img.shields.io/static/v1?label=Image_size&message=39.8MB&color=informational&logo=Docker&link=https%3A%2F%2Fhub.docker.com%2Frepository%2Fdocker%2Funitrade23%2Fkube-pull-exporter%2Fgeneral)

![image (12)](https://github.com/unitrade/kube-pull-exporter/assets/28563126/32991bc6-e121-4769-b5ee-0fc87991314e)

Kube Pull Exporter is a tool designed to monitor and export metrics from Kubernetes clusters. This repository contains the source code, documentation, and examples to help you get started with using Kube Pull Exporter in your projects.

## Getting Started
Kube Pull Exporter is built to work with Kubernetes clusters. Ensure you have a Kubernetes cluster up and running before proceeding with the installation and configuration steps.

## Installation
To install Kube Pull Exporter, follow these steps:

1. Clone the repository:
``` git clone https://github.com/unitrade/kube-pull-exporter.git ```
2. Navigate to the cloned directory:
``` cd kube-pull-exporter```
3. Install the necessary dependencies as per the project's documentation.

## Usage
To start monitoring your Kubernetes cluster with Kube Pull Exporter, run the following command from the root directory of the project:

```./kube-pull-exporter -namespaces=main,default -interval=1m```

## List of supported metrics
```
image_pull_duration_seconds_bucket
image_pull_duration_seconds_sum
image_pull_duration_seconds_count
```


## Example
```
# TYPE image_pull_duration_seconds histogram
image_pull_duration_seconds_bucket{image="nginx",le="1"} 0
image_pull_duration_seconds_bucket{image="nginx",le="2"} 0
image_pull_duration_seconds_bucket{image="nginx",le="5"} 1
image_pull_duration_seconds_bucket{image="nginx",le="10"} 1
image_pull_duration_seconds_bucket{image="nginx",le="15"} 1
image_pull_duration_seconds_bucket{image="nginx",le="20"} 1
image_pull_duration_seconds_bucket{image="nginx",le="30"} 1
image_pull_duration_seconds_bucket{image="nginx",le="+Inf"} 1
image_pull_duration_seconds_sum{image="nginx"} 2.213971983
image_pull_duration_seconds_count{image="nginx"} 1
```
