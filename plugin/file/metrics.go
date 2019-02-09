package file

import (
	"github.com/coredns/coredns/plugin"

	"github.com/prometheus/client_golang/prometheus"
)

const subsystem = "secondary"

// Metrics for the secondary plugin
var (
	UpdatedSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: plugin.Namespace,
		Subsystem: subsystem,
		Name:      "update_duration_seconds",
		Buckets:   plugin.TimeBuckets,
		Help:      "Histogram of the duration of an update",
	}, []string{"zone"})

	UpdatedSerial = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: plugin.Namespace,
		Subsystem: subsystem,
		Name:      "serial",
		Help:      "Gauge of the last updated serial",
	}, []string{"zone"})

	UpdatedSuccess = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: plugin.Namespace,
		Subsystem: subsystem,
		Name:      "last_updated_seconds",
		Help:      "Gauge of the last updated timestamp",
	}, []string{"zone"})

	RefreshCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: subsystem,
		Name:      "refreshes_total",
		Help:      "Counter of the number of refreshes",
	}, []string{"zone"})

	RefreshSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: plugin.Namespace,
		Subsystem: subsystem,
		Name:      "refresh_seconds",
		Buckets:   plugin.TimeBuckets,
		Help:      "Histogram of the duration of a refresh",
	}, []string{"zone"})

	FailedRefreshCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: subsystem,
		Name:      "refreshes_failed_total",
		Help:      "Counter of the number of failed refresh attempts",
	}, []string{"zone"})
)
