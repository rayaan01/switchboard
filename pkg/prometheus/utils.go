package prometheus

import "github.com/prometheus/client_golang/prometheus"

var SetCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "set_requests",
		Help: "Number of set requests handled",
	},
)

var GetCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "get_requests",
		Help: "Number of get requests handled",
	},
)

var DelCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "del_requests",
		Help: "Number of del requests handled",
	},
)

var SetHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Namespace: "set_time",
	Name:      "set_time",
	Help:      "Time taken to execute set query on the key value store",
	Buckets:   []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
}, []string{"key"})
