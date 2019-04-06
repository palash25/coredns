package rewrite

import (
	"github.com/coredns/coredns/plugin"

	"github.com/prometheus/client_golang/prometheus"
)

// Metrics for the rewrite plugin
var (
	// rewrite_request_name_total{server, match_type, src, dest}
	RequestRewriteNameCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "rewrite",
		Name:      "request_rewrite_name_total",
		Help:      "Counter of request name rewrites",
	}, []string{"server", "match_type", "src", "dest"})

	// rewrite_request_name_noop_total{server, match_type, src}
	RequestRewriteNoopCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "rewrite",
		Name:      "request_rewrite_name_noop_total",
		Help:      "Counter of request name rewrites that resulted in noops",
	}, []string{"server", "match_type", "src", "dest"})

	// rewrite_request_type_total{server, src_type, dest_type}
	RequestRewriteTypeCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "rewrite",
		Name:      "request_rewrite_type_total",
		Help:      "Counter of request type rewrites",
	}, []string{"server", "src", "dest"})

	// rewrite_request_class_total{server, src_class, dest_class}
	RequestRewriteClassCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "rewrite",
		Name:      "request_rewrite_class_total",
		Help:      "Counter of request type rewrites",
	}, []string{"server", "src", "dest"})

	// rewrite_request_edns0_total{server}
	RequestRewriteEDNSCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "rewrite",
		Name:      "request_rewrite_edns_total",
		Help:      "Counter of request type rewrites",
	}, []string{"server", "type", "src", "dest"})

	// rewrite_response_ttl_total{server, match_type, ttl}
	ResponseRewriteTTLCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "rewrite",
		Name:      "rewrite_response_ttl_total",
		Help:      "Counter of response ttl rewrites",
	}, []string{"server", "match_type", "ttl"})

	// rewrite_response_name_total{server, src, dest}
	ResponseRewriteNameCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "rewrite",
		Name:      "rewrite_response_name_total",
		Help:      "Counter of response ttl rewrites",
	}, []string{"server", "src", "dest"})

	// rewrite_rules_total{server, field_type}
	// rewrite_enabled{server}
)
