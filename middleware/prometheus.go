package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

// @Summary      List all Prometheus metrics
// @Description  Provide a list of all currently known Prometheus metrics
// @Tags         metrics
// @Produce      plain
// @Success      200  {string}  string "Prometheus metrics line by line"
// @Router       /metrics [get]
func RegisterPrometheus(router *gin.Engine) {
	// get global Monitor object
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(router)
}
