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
