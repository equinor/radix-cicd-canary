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
