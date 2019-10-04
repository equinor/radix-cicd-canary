package test

import (
	"time"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	log "github.com/sirupsen/logrus"
)

var (
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
type Fn func(env env.Env, suiteName string) (success bool, err error)

// ResultFn Prototype of result of a test function (success or fail)
type ResultFn func(testName string)

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
	SuccessFn   ResultFn
	FailFn      ResultFn
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
	suiteName := suite.Name
	setupFailed := false
	start := time.Now()

	for _, setup := range suite.Setup {
		log.Info(setup.Description)
		success := runTest(env, setup, suiteName)
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
	suiteName := suite.Name
	start := time.Now()

	for _, test := range suite.Tests {
		log.Info(test.Description)
		success := runTest(env, test, suiteName)
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
	suiteName := suite.Name
	start := time.Now()

	log.Infof("Running teardown tests in suite %s", suite.Name)
	for _, test := range suite.Teardown {
		log.Info(test.Description)
		runTest(env, test, suiteName)
	}

	end := time.Now()
	elapsed := end.Sub(start)
	scenarioDuration[suite.Name] = scenarioDuration[suite.Name] + elapsed
}

func runTest(env env.Env, testToRun Spec, suiteName string) bool {
	start := time.Now()

	success, err := testToRun.Test(env, suiteName)
	if !success {
		testToRun.FailFn(testToRun.Name)
		log.Errorf("Error calling %s: %v", testToRun.Name, err)
	} else {
		testToRun.SuccessFn(testToRun.Name)
		log.Info("Test success")
	}

	end := time.Now()
	elapsed := end.Sub(start)

	addTestDuration(testToRun.Name, elapsed.Seconds())
	log.Infof("Elapsed time: %v", elapsed)
	return success
}

func addTestDuration(testname string, durationSec float64) {
	testDurations.With(prometheus.Labels{"testName": testname}).Add(durationSec)
}
