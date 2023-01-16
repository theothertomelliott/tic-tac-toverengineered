package opentelemetry

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func CreatePrometheusHandler() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
