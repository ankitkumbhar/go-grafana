package middleware

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	reqs         *prometheus.CounterVec
	ResponseTime prometheus.Summary
	TotalReqs    *prometheus.CounterVec
)

// RegisterCollector
func RegisterCollector() {
	TotalReqs = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:      "custom_requests_total",
		Help:      "Number of requests",
	}, []string{"path"})
	prometheus.Register(TotalReqs)

	ResponseTime = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:        "custom_response_time",
		Help:        "Request response time",
	})
	prometheus.Register(ResponseTime)
}

// NewPrometheus
func NewPrometheus(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(res, req)
		path:=req.URL.Path

		// ignore /metrics path
		if req.URL.Path != "/metrics" {
			TotalReqs.WithLabelValues(path).Inc()
			ResponseTime.Observe(time.Since(startTime).Seconds())
		}
	}
	return http.HandlerFunc(fn)
}


