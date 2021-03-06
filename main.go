package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	gib3700 = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gibibytes_3700",
		Help: "This value represents 3700 * 1024 * 1024 * 1024 bytes",
	}, []string{"label1"})
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "200 ok")
}

func main() {
	gib3700.WithLabelValues("please-print-me-as-3700gib").Set(3700)

	http.HandleFunc("/health", healthHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
