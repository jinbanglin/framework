package addtransport

import (
	"moss/metrics"
	"moss/metrics/prometheus"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	Counters       metrics.Counter
	SummarySuccess metrics.Histogram
	SummaryError   metrics.Histogram
}

func NewMetrics() *Metrics {
	mts := &Metrics{}
	mts.setSummaryLabelNames()
	mts.setSummary()
	mts.setCounters()
	return mts
}

func (m *Metrics) setCounters() {
	m.Counters = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "group",
		Subsystem: "Invoking",
		Name:      "protocol_counter",
		Help:      "Protocol code of requests received",
	}, []string{"method", "error"})
}

func (m *Metrics) setSummary() {
	m.SummarySuccess = prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "group",
		Subsystem: "Invoking",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, []string{"method", "error"})
}

func (m *Metrics) setSummaryLabelNames() {
	m.SummaryError = prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "group",
		Subsystem: "Invoking",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{"method", "error"})
}
