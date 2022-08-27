package middleware

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// @Summary      Get Prometheus metrics
// @Description  Provide a list of all currently provided metrics
// @Tags         metrics
// @Produce      plain
// @Success      200  {string}  string "metrics line by line"
// @Router       /metrics [get]
func RegisterPrometheus() {
	//register prometheus
	http.Handle("/metrics", promhttp.Handler())

}
