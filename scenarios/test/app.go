package test

import (
	"time"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"

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
type Fn func(env env.Env) (success bool, err error)

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

// Runner Instance
type Runner struct {
	env env.Env
}

// NewRunner Constructor
func NewRunner(env env.Env) Runner {
	return Runner{
		env,
	}
}

// Run the suite
func (runner Runner) Run(suites ...Suite) {
	setupFailed := false
	scenarioDuration := make(map[string]time.Duration)

	// Run all suite setup
	for _, suite := range suites {
		setupFailed = runSuiteSetup(runner.env, suite, scenarioDuration)
		if setupFailed {
			break
		}
	}

	if !setupFailed {
		for _, suite := range suites {
			runSuiteTests(runner.env, suite, scenarioDuration)
		}
	}

	// Run all suite teardown
	for _, suite := range suites {
		runSuiteTeardown(runner.env, suite, scenarioDuration)
	}

	for scenario, elapsed := range scenarioDuration {
		scenarioDurations.With(prometheus.Labels{"scenario": scenario}).Add(elapsed.Seconds())
		log.Infof("%s elapsed time: %v", scenario, elapsed)
	}
}

func runSuiteSetup(env env.Env, suite Suite, scenarioDuration map[string]time.Duration) bool {
	setupFailed := false
	start := time.Now()

	for _, setup := range suite.Setup {
		log.Info(setup.Description)
		success := runTest(env, setup)
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

func runSuiteTests(env env.Env, suite Suite, scenarioDuration map[string]time.Duration) {
	start := time.Now()

	for _, test := range suite.Tests {
		log.Info(test.Description)
		success := runTest(env, test)
		if !success {
			log.Warnf("Test %s fail. Will escape remaining tests in suite %s", test.Name, suite.Name)
			break
		}
	}

	end := time.Now()
	elapsed := end.Sub(start)
	scenarioDuration[suite.Name] = scenarioDuration[suite.Name] + elapsed
}

func runSuiteTeardown(env env.Env, suite Suite, scenarioDuration map[string]time.Duration) {
	start := time.Now()

	log.Infof("Running teardown tests in suite %s", suite.Name)
	for _, test := range suite.Teardown {
		log.Info(test.Description)
		runTest(env, test)
	}

	end := time.Now()
	elapsed := end.Sub(start)
	scenarioDuration[suite.Name] = scenarioDuration[suite.Name] + elapsed
}

func runTest(env env.Env, testToRun Spec) bool {
	start := time.Now()

	success, err := testToRun.Test(env)
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
