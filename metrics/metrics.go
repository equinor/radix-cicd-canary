package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	testDurations = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radix_test_duration",
			Help: "Duration of test",
		},
		[]string{"testName"},
	)
	scenarioDurations = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radix_scenario_duration",
			Help: "Duration of Scenario",
		},
		[]string{"scenario"},
	)
)

// AddScenarioDuration Adds scenario duration for provided scenario
func AddScenarioDuration(scenario string, elapsed time.Duration) {
	scenarioDurations.With(prometheus.Labels{"scenario": scenario}).Set(elapsed.Seconds())
}

// AddTestDuration Adds duration for provided test
func AddTestDuration(testname string, elapsed time.Duration) {
	testDurations.With(prometheus.Labels{"testName": testname}).Set(elapsed.Seconds())
}

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

// AddTestSuccess adds 1 to the success counter metrics for provided test
func AddTestSuccess(testname string) {
	success.With(prometheus.Labels{"testName": testname}).Add(1)
}

// AddTestNoSuccess adds 0 to the success counter metrics for provided test
func AddTestNoSuccess(testname string) {
	success.With(prometheus.Labels{"testName": testname}).Add(0)
}

// AddTestError adds 1 to the errors counter metrics for provided test
func AddTestError(testname string) {
	errors.With(prometheus.Labels{"testName": testname}).Add(1)
}

// AddTestNoError adds 0 to the errors counter metrics for provided test
func AddTestNoError(testname string) {
	errors.With(prometheus.Labels{"testName": testname}).Add(0)
}
