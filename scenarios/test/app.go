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
	Setup    []Spec
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
func Run(suites ...Suite) {
	setupFailed := false
	scenarioDuration := make(map[string]time.Duration)

	// Run all suite setup
	for _, suite := range suites {
		setupFailed = runSuiteSetup(suite, scenarioDuration)
		if setupFailed {
			break
		}
	}

	if !setupFailed {
		for _, suite := range suites {
			runSuiteTests(suite, scenarioDuration)
		}
	}

	// Run all suite teardown
	for _, suite := range suites {
		runSuiteTeardown(suite, scenarioDuration)
	}

	for scenario, elapsed := range scenarioDuration {
		scenarioDurations.With(prometheus.Labels{"scenario": scenario}).Add(elapsed.Seconds())
		log.Infof("%s elapsed time: %v", scenario, elapsed)
	}
}

func runSuiteSetup(suite Suite, scenarioDuration map[string]time.Duration) bool {
	setupFailed := false
	start := time.Now()

	for _, setup := range suite.Setup {
		log.Info(setup.Description)
		success := runTest(setup)
		if !success {
			setupFailed = true
			log.Warnf("Setup %s fail in suite %s. Will escape tests, and just run teardowns", setup.Name, suite.Name)
			break
		}
	}

	end := time.Now()
	elapsed := end.Sub(start)
	scenarioDuration[suite.Name] = elapsed
	return setupFailed
}

func runSuiteTests(suite Suite, scenarioDuration map[string]time.Duration) {
	start := time.Now()

	for _, test := range suite.Tests {
		log.Info(test.Description)
		success := runTest(test)
		if !success {
			log.Warnf("Test %s fail. Will escape remaining tests in suite %s", test.Name, suite.Name)
			break
		}
	}

	end := time.Now()
	elapsed := end.Sub(start)
	scenarioDuration[suite.Name] = scenarioDuration[suite.Name] + elapsed
}

func runSuiteTeardown(suite Suite, scenarioDuration map[string]time.Duration) {
	start := time.Now()

	log.Infof("Running teardown tests in suite %s", suite.Name)
	for _, test := range suite.Teardown {
		log.Info(test.Description)
		runTest(test)
	}

	end := time.Now()
	elapsed := end.Sub(start)
	scenarioDuration[suite.Name] = scenarioDuration[suite.Name] + elapsed
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
