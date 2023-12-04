package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const MetricsNamespace = "prom_agg_gateway"

var PromRegistry = prometheus.NewRegistry()

var metricsClock func() time.Time

func init() {
	// NOTE(lxf): Clock is overridable for testing purposes.
	metricsClock = time.Now

	PromRegistry.MustRegister(
		TotalFamiliesGauge,
		MetricCountByFamily,
		MetricPushes,
	)
}

var TotalFamiliesGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: MetricsNamespace,
		Name:      "total_families",
		Help:      "Total number of metric families",
	},
)

var MetricCountByFamily = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: MetricsNamespace,
		Name:      "metrics_by_family",
		Help:      "Metric count by family",
	},
	[]string{
		"family",
	},
)

var MetricCountByType = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: MetricsNamespace,
		Name:      "metrics_by_type",
		Help:      "Metric count by type",
	},
	[]string{
		"metric_type",
	},
)

var MetricPushes = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: MetricsNamespace,
		Name:      "metric_pushes",
		Help:      "Total number of metric push requests, per job",
	},
	[]string{
		"push_job",
	},
)
