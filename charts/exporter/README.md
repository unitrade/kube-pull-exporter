## Usage

## Prerequisites

* Install the follow packages: ``git``, ``kubectl``, ``helm``, ``helm-docs``.

## How to install

Access a Kubernetes cluster.

Add a chart helm repository with follow commands:

```console
helm repo add kube-pull-exporter https://unitrade.github.io/kube-pull-exporter/   

helm repo update
```

List versions of ``kube-pull-exporter`` chart available to installation:

```console
helm search repo kube-pull-exporter -l
```
Test the installation with command:

```console
helm install kube-pull-exporter kube-pull-exporter -f values.yaml -n NAMESPACE --debug --dry-run
```

Install chart with command:

```console
helm install kube-pull-exporter kube-pull-exporter -f values.yaml -n NAMESPACE
```

Get the pods lists by running this commands:

```console
kubectl get pods -A | grep 'kube-pull-exporter'
```

## How to uninstall

Remove application with command.

```console
helm uninstall kube-pull-exporter -n NAMESPACE
```