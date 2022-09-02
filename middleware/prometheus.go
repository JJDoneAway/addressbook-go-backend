package middleware

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	prometheusHttp "github.com/prometheus/client_golang/prometheus/promhttp"
)

// Logger is a middleware handler that does request logging
type Prometheus struct {
	handler http.Handler
}

// wrapper for the response writer to get the http status code
type statusRecorder struct {
	http.ResponseWriter
	status int
}

var (
	rpcCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        "http_counter",
			Help:        "Number of HTTP requests.",
			ConstLabels: map[string]string{},
		},
		[]string{"path", "method", "status"},
	)

	// Create a summary to track fictional interservice RPC latencies for three
	// distinct services with different latency distributions. These services are
	// differentiated via a "service" label.
	rpcDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "http_durations_seconds",
			Help:       "Latency distributions of HTTP requests.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"path", "method", "status"},
	)
)

// NewPrometheus constructs a new Prometheus middleware handler
func NewPrometheus(handlerToWrap http.Handler) *Prometheus {
	return &Prometheus{handlerToWrap}
}

// Middleware wrapper around every http call to do some prometheus metric exposure
// in here
func (prometheus *Prometheus) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	rec := statusRecorder{w, http.StatusOK}
	prometheus.handler.ServeHTTP(&rec, r)

	pathFragments := regexp.MustCompile(`^/([^/]+)/?`).FindStringSubmatch(r.URL.Path)
	path := "/"
	if len(pathFragments) > 0 {
		path = pathFragments[1]
	}
	rpcCounter.WithLabelValues(path, r.Method, strconv.Itoa(rec.status)).Inc()
	rpcDurations.WithLabelValues(path, r.Method, strconv.Itoa(rec.status)).Observe(float64(time.Since(start).Seconds()))

}

// wrapper around the writer func itself
func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

// @Summary      Get Prometheus metrics
// @Description  Provide a list of all currently provided metrics
// @Tags         metrics
// @Produce      plain
// @Success      200  {string}  string "metrics line by line"
// @Router       /metrics [get]
func RegisterPrometheus(mux *http.ServeMux) {
	prometheus.MustRegister(rpcCounter)
	prometheus.MustRegister(rpcDurations)
	// Add Go module build info.
	prometheus.MustRegister(collectors.NewBuildInfoCollector())

	// Expose the registered metrics via HTTP.
	mux.Handle("/metrics", prometheusHttp.HandlerFor(
		prometheus.DefaultGatherer,
		prometheusHttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))
}
