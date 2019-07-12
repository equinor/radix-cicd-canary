package happypath

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	log "github.com/sirupsen/logrus"
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
	testDurations = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "radix_test_duration",
			Help: "Duration of test",
		},
		[]string{"testName"},
	)
	scenarioDurations = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "radix_scenario_duration",
			Help: "Duration of Scenario",
		},
		[]string{"scenario"},
	)
)

type testFn func() (testname string)

// Run the happypath scenario
func Run() {
	start := time.Now()

	log.Infof("List applications")
	runTest(listApplications)
	// log.Infof("Register application")
	// registerApplication()
	// log.Infof("Register application with no deploy key")
	// registerApplicationWithNoDeployKey()
	// log.Infof("Build application")
	// buildApplication()
	// log.Infof("Set secret")
	// setSecret()
	// log.Infof("Check alias responding")
	// defaultAliasResponding()
	// log.Infof("Check access to application user should not be able to access")
	// unauthorizedAccess()
	// log.Infof("Delete applications")
	// runTest(deleteApplications)

	end := time.Now()
	elapsed := end.Sub(start)

	scenarioDurations.With(prometheus.Labels{"scenario": "Happy Path"}).Add(elapsed.Seconds())
	log.Infof("Happy path elapsed time: %v", elapsed)
}

func runTest(testToRun testFn) {
	start := time.Now()

	testName := testToRun()

	end := time.Now()
	elapsed := end.Sub(start)

	addTestDuration(testName, elapsed.Seconds())
	log.Infof("Elapsed time: %v", elapsed)

}

func addTestSuccess(testname string) {
	success.With(prometheus.Labels{"testName": testname}).Add(1)
}

func addTestError(testname string) {
	errors.With(prometheus.Labels{"testName": testname}).Add(1)
}

func addTestDuration(testname string, durationSec float64) {
	testDurations.With(prometheus.Labels{"testName": testname}).Add(durationSec)
}

func getApp1Name() string {
	return "an_app"
}

func getApp2Name() string {
	return "a_second_app"
}
