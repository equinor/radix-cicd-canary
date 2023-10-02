package test

import (
	"time"

	"github.com/equinor/radix-cicd-canary/metrics"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/rs/zerolog/log"
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
		log.Info().Str("scenario", scenario).Dur("elapsed", elapsed).Msgf("elapsed time: %v", elapsed)
	}
}

func runSuiteSetup(cfg config.Config, suite Suite, scenarioDuration map[string]time.Duration) bool {
	suiteName := suite.Name
	setupFailed := false
	start := time.Now()
	logger := log.With().Str("suite", suiteName).Logger()
	logger.Debug().Str("suite", suiteName).Msgf("Setting-up suite '%s'", suiteName)

	for _, setup := range suite.Setup {
		logger.Info().Msg(setup.Description)
		success := runTest(cfg, setup, suiteName)
		if !success {
			setupFailed = true
			logger.Error().Str("suite", suite.Name).Str("setupname", setup.Name).Msgf("!!!!!!!!!!!!!!!!!!!!!!!!! Setup %s fail in suite %s. Will escape tests, and just run teardowns !!!!!!!!!!!!!!!!!!!!!!!!!", setup.Name, suite.Name)
			break
		}
		logger.Debug().Msgf("Setup success %s", setup.Description)
	}

	end := time.Now()
	elapsed := end.Sub(start)
	scenarioDuration[suite.Name] = elapsed
	return setupFailed
}

func runSuiteTests(cfg config.Config, suite Suite, scenarioDuration map[string]time.Duration) {
	suiteName := suite.Name
	start := time.Now()
	logger := log.With().Str("suite", suiteName).Logger()

	for _, test := range suite.Tests {
		logger.Info().Msg(test.Description)
		success := runTest(cfg, test, suiteName)
		if !success {
			logger.Warn().Str("test", test.Name).Msgf("!!!!!!!!!!!!!!!!!!!!!!!!! Test %s fail. Will escape remaining tests in the suite !!!!!!!!!!!!!!!!!!!!!!!!!!!", test.Name)
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
	logger := log.With().Str("suite", suiteName).Logger()

	logger.Debug().Msg("Running teardown tests in suite")
	for _, test := range suite.Teardown {
		logger.Info().Msg(test.Description)
		runTest(cfg, test, suiteName)
	}
	logger.Debug().Msg("Teardown complete")

	end := time.Now()
	elapsed := end.Sub(start)
	scenarioDuration[suite.Name] = scenarioDuration[suite.Name] + elapsed
}

func runTest(cfg config.Config, testToRun Spec, suiteName string) bool {
	start := time.Now()
	logger := log.With().Str("suite", suiteName).Str("test", testToRun.Name).Logger()

	logger.Debug().Msg("Running test")

	err := testToRun.Test(cfg, suiteName)
	if err != nil {
		testToRun.FailFn(testToRun.Name)
		logger.Error().Err(err).Msg("Test failed")
	} else {
		testToRun.SuccessFn(testToRun.Name)
		logger.Debug().Msg("Test success")
	}

	end := time.Now()
	elapsed := end.Sub(start)

	metrics.AddTestDuration(testToRun.Name, elapsed)
	log.Info().Dur("elapsed", elapsed).Msgf("elapsed time: %v", elapsed)
	return err == nil
}
