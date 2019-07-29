package test

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

// Fn Prototype of a test function
type Fn func() (success bool, err error)

// Suite Holds a list of tests
type Suite struct {
	Name     string
	Tests    []Spec
	Teardown []Spec
}

// Spec Describes a single test
type Spec struct {
	Name        string
	Description string
	Test        Fn
}

// Run the suite
func Run(suite Suite) {
	start := time.Now()

	for _, test := range suite.Tests {
		log.Info(test.Description)
		success := runTest(test)
		if !success {
			log.Warnf("Test %s fail. Will escape remaining tests", test.Name)
			break
		}
	}

	log.Info("Running teardown tests")
	for _, test := range suite.Teardown {
		log.Info(test.Description)
		runTest(test)
	}

	end := time.Now()
	elapsed := end.Sub(start)

	scenarioDurations.With(prometheus.Labels{"scenario": suite.Name}).Add(elapsed.Seconds())
	log.Infof("%s elapsed time: %v", suite.Name, elapsed)
}

func runTest(testToRun Spec) bool {
	start := time.Now()

	success, err := testToRun.Test()
	if !success {
		addTestNoSuccess(testToRun.Name)
		addTestError(testToRun.Name)
		log.Errorf("Error calling %s: %v", testToRun.Name, err)
	} else {
		addTestSuccess(testToRun.Name)
		addTestNoError(testToRun.Name)
		log.Info("Test success")
	}

	end := time.Now()
	elapsed := end.Sub(start)

	addTestDuration(testToRun.Name, elapsed.Seconds())
	log.Infof("Elapsed time: %v", elapsed)
	return success
}

func addTestSuccess(testname string) {
	success.With(prometheus.Labels{"testName": testname}).Add(1)
}

func addTestNoSuccess(testname string) {
	success.With(prometheus.Labels{"testName": testname}).Add(0)
}

func addTestError(testname string) {
	errors.With(prometheus.Labels{"testName": testname}).Add(1)
}

func addTestNoError(testname string) {
	errors.With(prometheus.Labels{"testName": testname}).Add(0)
}

func addTestDuration(testname string, durationSec float64) {
	testDurations.With(prometheus.Labels{"testName": testname}).Add(durationSec)
}
