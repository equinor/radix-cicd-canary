package test

import (
	"context"
	"time"

	"github.com/equinor/radix-cicd-canary/metrics"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/rs/zerolog/log"
)

// Fn Prototype of a test function
type Fn func(ctx context.Context, cfg config.Config) error

// ResultFn Prototype of result of a test function (success or fail)
type ResultFn func(ctx context.Context, testName string)

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
func (runner Runner) Run(ctx context.Context, suites ...Suite) {
	setupFailed := false
	scenarioDuration := make(map[string]time.Duration)

	// Run all suite setup
	for _, suite := range suites {
		setupCtx := log.Ctx(ctx).With().
			Str("stage", "setup").
			Str("suite", suite.Name).
			Logger().WithContext(ctx)

		setupFailed = runSuiteSetup(setupCtx, runner.cfg, suite, scenarioDuration)
		if setupFailed {
			break
		}
	}

	if !setupFailed {
		for _, suite := range suites {
			testCtx := log.Ctx(ctx).With().
				Str("stage", "test").
				Str("suite", suite.Name).
				Logger().WithContext(ctx)

			runSuiteTests(testCtx, runner.cfg, suite, scenarioDuration)
		}
	}

	// Run all suite teardown
	for _, suite := range suites {
		teardownCtx := log.Ctx(ctx).With().
			Str("stage", "teardown").
			Str("suite", suite.Name).
			Logger().WithContext(ctx)
		runSuiteTeardown(teardownCtx, runner.cfg, suite, scenarioDuration)
	}

	for scenario, elapsed := range scenarioDuration {
		metrics.AddScenarioDuration(scenario, elapsed)
		log.Ctx(ctx).Info().Str("scenario", scenario).Dur("elapsed", elapsed).Msgf("elapsed time: %v", elapsed)
	}
}

func runSuiteSetup(ctx context.Context, cfg config.Config, suite Suite, scenarioDuration map[string]time.Duration) bool {
	setupFailed := false
	start := time.Now()
	log.Ctx(ctx).Debug().Msg("Setting-up suite")

	for _, setup := range suite.Setup {
		testCtx := log.Ctx(ctx).With().Str("test", setup.Name).Logger().WithContext(ctx)
		log.Ctx(testCtx).Info().Msg(setup.Description)
		success := runTest(testCtx, cfg, setup)
		if !success {
			setupFailed = true
			log.Ctx(ctx).Error().Str("setup", setup.Name).Msgf("!!!!!!!!!!!!!!!!!!!!!!!!! Setup %s fail in suite %s. Will escape tests, and just run teardowns !!!!!!!!!!!!!!!!!!!!!!!!!", setup.Name, suite.Name)
			break
		}
		log.Ctx(ctx).Debug().Msgf("Setup success %s", setup.Description)
	}

	end := time.Now()
	elapsed := end.Sub(start)
	scenarioDuration[suite.Name] = elapsed
	return setupFailed
}

func runSuiteTests(ctx context.Context, cfg config.Config, suite Suite, scenarioDuration map[string]time.Duration) {
	start := time.Now()

	for _, test := range suite.Tests {
		testCtx := log.Ctx(ctx).With().Str("test", test.Name).Logger().WithContext(ctx)
		log.Ctx(testCtx).Info().Msg(test.Description)

		success := runTest(testCtx, cfg, test)
		if !success {
			log.Ctx(testCtx).Warn().Msgf("!!!!!!!!!!!!!!!!!!!!!!!!! Test %s fail. Will escape remaining tests in the suite !!!!!!!!!!!!!!!!!!!!!!!!!!!", test.Name)
			break
		}
	}

	end := time.Now()
	elapsed := end.Sub(start)
	scenarioDuration[suite.Name] = scenarioDuration[suite.Name] + elapsed
}

func runSuiteTeardown(ctx context.Context, cfg config.Config, suite Suite, scenarioDuration map[string]time.Duration) {
	start := time.Now()

	log.Ctx(ctx).Debug().Msg("Running teardown tests in suite")
	for _, test := range suite.Teardown {
		testCtx := log.Ctx(ctx).With().Str("test", test.Name).Logger().WithContext(ctx)
		log.Ctx(testCtx).Info().Msg(test.Description)
		runTest(testCtx, cfg, test)
	}
	log.Ctx(ctx).Debug().Msg("Teardown complete")

	end := time.Now()
	elapsed := end.Sub(start)
	scenarioDuration[suite.Name] = scenarioDuration[suite.Name] + elapsed
}

func runTest(ctx context.Context, cfg config.Config, testToRun Spec) bool {
	start := time.Now()

	log.Ctx(ctx).Debug().Msg("Running test")

	err := testToRun.Test(ctx, cfg)
	if err != nil {
		testToRun.FailFn(ctx, testToRun.Name)
		log.Ctx(ctx).Error().Err(err).Msg("Test failed")
	} else {
		testToRun.SuccessFn(ctx, testToRun.Name)
		log.Ctx(ctx).Debug().Msg("Test success")
	}

	end := time.Now()
	elapsed := end.Sub(start)

	metrics.AddTestDuration(testToRun.Name, elapsed)
	log.Ctx(ctx).Info().Dur("elapsed", elapsed).Msgf("elapsed time: %v", elapsed)
	return err == nil
}
