package service

import (
	"context"
	"errors"

	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger zerolog.Logger

// Reach tests that we are able to reach radix-canary-golang-prod endpoint
func Reach(ctx context.Context, cfg config.Config, suiteName string) error {

	client := httpUtils.GetHTTPDefaultClient()
	url := "http://www.radix-canary-golang-prod:5000/health"
	log.Ctx(ctx).Info().Str("url", url).Msg("requesting data")

	// Run tests service
	_, err := client.Get(url)
	if err == nil {
		// Failing test
		return errors.New("no error was returned when attempting to access service")
	}

	// Successful
	return nil
}

// Success is a function after a call to Reach succeeds
func Success(testName string) {
	nspMetrics.AddServiceUnreachable()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	logger.Info().Str("test", testName).Msg("Test: SUCCESS")
}

// Fail is a function after a call to Reach failed
func Fail(testName string) {
	nspMetrics.AddServiceReachable()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	logger.Info().Str("test", testName).Msg("Test: FAIL")
}
