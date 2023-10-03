package service

import (
	"context"
	"errors"

	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/rs/zerolog/log"
)

// Reach tests that we are able to reach radix-canary-golang-prod endpoint
func Reach(ctx context.Context, cfg config.Config) error {

	client := httpUtils.GetHTTPDefaultClient()
	url := "http://www.radix-canary-golang-prod:5000/health"
	log.Ctx(ctx).Debug().Msgf("requesting data from url %s", url)

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
func Success(ctx context.Context, testName string) {
	nspMetrics.AddServiceUnreachable()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: SUCCESS")
}

// Fail is a function after a call to Reach failed
func Fail(ctx context.Context, testName string) {
	nspMetrics.AddServiceReachable()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: FAIL")
}
