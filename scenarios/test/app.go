package test

import (
	"time"

	"github.com/equinor/radix-cicd-canary/metrics"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	log "github.com/sirupsen/logrus"
)

// Fn Prototype of a test function
type Fn func(cfg config.Config, suiteName string) error

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
	cfg config.Config
}

// NewRunner Constructor
func NewRunner(cfg config.Config) Runner {
	return Runner{
		cfg,
	}
}

// Run the suite
func (runner Runner) Run(suites ...Suite) {
	setupFailed := false
	scenarioDuration := make(map[string]time.Duration)

	// Run all suite setup
	for _, suite := range suites {
		setupFailed = runSuiteSetup(runner.cfg, suite, scenarioDuration)
		if setupFailed {
			break
		}
	}

	if !setupFailed {
		for _, suite := range suites {
			runSuiteTests(runner.cfg, suite, scenarioDuration)
		}
	}

	// Run all suite teardown
	for _, suite := range suites {
		runSuiteTeardown(runner.cfg, suite, scenarioDuration)
	}

	for scenario, elapsed := range scenarioDuration {
		metrics.AddScenarioDuration(scenario, elapsed)
		log.Infof("%s elapsed time: %v", scenario, elapsed)
	}
}

func runSuiteSetup(cfg config.Config, suite Suite, scenarioDuration map[string]time.Duration) bool {
	suiteName := suite.Name
	setupFailed := false
	start := time.Now()
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	logger.Debugf("Setting-up suite '%s'", suiteName)

	for _, setup := range suite.Setup {
		logger.Info(setup.Description)
		success := runTest(cfg, setup, suiteName)
		if !success {
			setupFailed = true
			logger.Errorf("!!!!!!!!!!!!!!!!!!!!!!!!! Setup %s fail in suite %s. Will escape tests, and just run teardowns !!!!!!!!!!!!!!!!!!!!!!!!!", setup.Name, suite.Name)
			break
		}
		logger.Debugf("Setup success %s", setup.Description)
	}

	end := time.Now()
	elapsed := end.Sub(start)
	scenarioDuration[suite.Name] = elapsed
	return setupFailed
}

func runSuiteTests(cfg config.Config, suite Suite, scenarioDuration map[string]time.Duration) {
	suiteName := suite.Name
	start := time.Now()
	logger := log.WithFields(log.Fields{"Suite": suiteName})

	for _, test := range suite.Tests {
		logger.Info(test.Description)
		success := runTest(cfg, test, suiteName)
		if !success {
			logger.Warnf("!!!!!!!!!!!!!!!!!!!!!!!!! Test %s fail. Will escape remaining tests in the suite !!!!!!!!!!!!!!!!!!!!!!!!!!!", test.Name)
			break
		}
	}

	end := time.Now()
	elapsed := end.Sub(start)
	scenarioDuration[suite.Name] = scenarioDuration[suite.Name] + elapsed
}

func runSuiteTeardown(cfg config.Config, suite Suite, scenarioDuration map[string]time.Duration) {
	suiteName := suite.Name
	start := time.Now()
	logger := log.WithFields(log.Fields{"Suite": suiteName})

	logger.Debugf("Running teardown tests in suite %s", suite.Name)
	for _, test := range suite.Teardown {
		logger.Info(test.Description)
		runTest(cfg, test, suiteName)
	}
	logger.Debug("Teardown complete")

	end := time.Now()
	elapsed := end.Sub(start)
	scenarioDuration[suite.Name] = scenarioDuration[suite.Name] + elapsed
}

func runTest(cfg config.Config, testToRun Spec, suiteName string) bool {
	start := time.Now()
	logger := log.WithFields(log.Fields{"Suite": suiteName})

	logger.Debugf("Running test '%s'", testToRun.Name)

	err := testToRun.Test(cfg, suiteName)
	if err != nil {
		testToRun.FailFn(testToRun.Name)
		logger.Errorf("Error calling %s: %v", testToRun.Name, err)
	} else {
		testToRun.SuccessFn(testToRun.Name)
		logger.Debug("Test success")
	}

	end := time.Now()
	elapsed := end.Sub(start)

	metrics.AddTestDuration(testToRun.Name, elapsed)
	logger.Infof("Elapsed time: %v", elapsed)
	return err == nil
}
