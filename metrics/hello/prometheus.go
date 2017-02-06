package metrics

import (
	kmetrics "github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// Testfunc implements asdas
func Testfunc() {
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{}, fieldKeys)

	requestLatency := kmetrics.NewTimer(kitprometheus.NewHistogramFrom(stdprometheus.HistogramOpts{}, fieldKeys))

	requestCount.Add(1)
	requestLatency.ObserveDuration()
}
