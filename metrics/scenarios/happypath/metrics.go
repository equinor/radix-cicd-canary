package happypath

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	errors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "radix_test_errors",
			Help: "Test errors",
		},
		[]string{"testName"},
	)
	success = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "radix_test_success",
			Help: "Test success",
		},
		[]string{"testName"},
	)
)

// AddTestSuccess adds 1 to the success counter metrics for happypath scenario
func AddTestSuccess(testname string) {
	success.With(prometheus.Labels{"testName": testname}).Add(1)
}

// AddTestNoSuccess adds 0 to the success counter metrics for happypath scenario
func AddTestNoSuccess(testname string) {
	success.With(prometheus.Labels{"testName": testname}).Add(0)
}

// AddTestError adds 1 to the errors counter metrics for happypath scenario
func AddTestError(testname string) {
	errors.With(prometheus.Labels{"testName": testname}).Add(1)
}

// AddTestNoError adds 0 to the errors counter metrics for happypath scenario
func AddTestNoError(testname string) {
	errors.With(prometheus.Labels{"testName": testname}).Add(0)
}
