package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

const (
	ImagePullDurationSecondsName = "image_pull_duration_seconds"
)

type Metrics struct {
	ImagePullDurationSecondsGauge *prometheus.GaugeVec
}

func (Metrics) Handler() http.Handler {
	return promhttp.Handler()
}

func New() *Metrics {
	ms := Metrics{
		ImagePullDurationSecondsGauge: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: ImagePullDurationSecondsName,
			Help: "Image pull duration",
		}, []string{"namespace", "pod_name", "image_name"}),
	}
	prometheus.MustRegister(
		ms.ImagePullDurationSecondsGauge,
	)
	return &ms
}
