package prometheus

import "github.com/prometheus/client_golang/prometheus"

// Metrics all prometheus metrics
type Metrics struct {
	NodeRequests prometheus.Histogram
}

// NewMetrics Creates new metrics holder
func NewMetrics() *Metrics {
	return &Metrics{
		NodeRequests: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "provider",
				Name:      "node_histogram",
				Help:      "Node latency",
			}),
	}
}