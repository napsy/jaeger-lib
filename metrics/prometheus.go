package metrics

import (
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	kmetrics "github.com/go-kit/kit/metrics"
)

func fuck() {
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{}, fieldKeys)

	requestLatency := kmetrics.NewTimer(kitprometheus.NewHistogramFrom(stdprometheus.HistogramOpts{}, fieldKeys))

	requestCount.Add(1)
	requestLatency.ObserveDuration()
}
