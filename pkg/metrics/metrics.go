package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

const (
	ImagePullDurationSecondsBucketName = "image_pull_duration_seconds"
)

type Metrics struct {
	ImagePullDurationSecondsHistogram *prometheus.HistogramVec
}

func (Metrics) Handler() http.Handler {
	return promhttp.Handler()
}

func New() *Metrics {
	ms := Metrics{
		ImagePullDurationSecondsHistogram: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    ImagePullDurationSecondsBucketName,
			Help:    "Duration of image pull operations.",
			Buckets: []float64{1, 2, 5, 10, 15, 20, 30},
		}, []string{"image"}),
	}
	prometheus.MustRegister(
		ms.ImagePullDurationSecondsHistogram,
	)
	return &ms
}
