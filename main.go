package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/equinor/radix-cicd-canary/scenarios/deployonly"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath"
	"github.com/equinor/radix-cicd-canary/scenarios/nsp"
	nsplong "github.com/equinor/radix-cicd-canary/scenarios/nsp-long"
	"github.com/equinor/radix-cicd-canary/scenarios/test"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

func init() {
	// If you get GOAWAY calling API with token using:
	// az account get-access-token
	// ...enable this line
	// os.Setenv("GODEBUG", "http2server=0,http2client=0")
}

func main() {
	cfg := config.NewConfig()
	ctx := context.Background()

	logLevel := cfg.GetLogLevel()
	pretty := cfg.GetPrettyPrint()
	zerolog.SetGlobalLevel(logLevel)
	zerolog.DurationFieldInteger = true
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	if pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly})
	}
	ctx = log.Logger.WithContext(ctx)

	log.Info().Msg("Starting...")
	log.Info().Msgf("Log level: %s", logLevel.String())

	sleepInterval := cfg.GetSleepIntervalBetweenTestRuns()
	happyPathSuite := happypath.TestSuite()
	deployOnlySuite := deployonly.TestSuite()

	nspSleepInterval := cfg.GetNSPSleepInterval()
	nspLongSleepInterval := cfg.GetNSPLongSleepInterval()
	nspSuite := nsp.TestSuite()
	nspLongSuite := nsplong.TestSuite()

	go runSuites(cfg, ctx, sleepInterval, happyPathSuite)
	go runSuites(cfg, ctx, sleepInterval, deployOnlySuite)
	go runSuites(cfg, ctx, nspSleepInterval, nspSuite)
	go runSuites(cfg, ctx, nspLongSleepInterval, nspLongSuite)

	log.Info().Msg("Started suites. Start metrics service.")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("Failed to listen")
		return
	}
	log.Info().Msg("Complete.")
}

func runSuites(environmentVariables config.Config, ctx context.Context, sleepInterval time.Duration, suites ...test.Suite) {
	log.Debug().Int("suites", len(suites)).Msg("Prepare to run suite(s)")
	suites = filterSuites(suites, environmentVariables)
	if len(suites) == 0 {
		log.Debug().Msg("No suites to run")
		return
	}

	log.Debug().Int("suites", len(suites)).Msg("Run suite(s)")
	runner := test.NewRunner(environmentVariables)
	for {
		runner.Run(ctx, suites...)
		time.Sleep(sleepInterval)
	}
}

func filterSuites(suites []test.Suite, environmentVariables config.Config) []test.Suite {
	filter := environmentVariables.GetSuiteList()
	if len(filter) == 0 {
		return suites
	}

	log.Debug().Msg("Filtering suites...")
	suitesToRun := make([]test.Suite, 0)
	isBlacklist := environmentVariables.GetSuiteListIsBlacklist()
	for _, suite := range suites {
		// pass the filter if mentioned and !isBlacklist OR if !mentioned and isBlacklist
		if contains(filter, suite.Name) != isBlacklist {
			log.Debug().Str("name", suite.Name).Msg("run suite")
			suitesToRun = append(suitesToRun, suite)
		} else {
			log.Debug().Str("name", suite.Name).Msg("skip suite")
		}
	}
	return suitesToRun
}

func contains(list []string, target string) bool {
	for _, item := range list {
		if target == item {
			return true
		}
	}
	return false
}
