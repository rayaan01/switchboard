package prometheus

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsListener() {
	prometheus.MustRegister(SetCounter)
	prometheus.MustRegister(GetCounter)
	prometheus.MustRegister(DelCounter)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		fmt.Println("Metrics listening on :8090")
		http.ListenAndServe(":8090", nil)
	}()
}
