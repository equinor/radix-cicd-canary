package ingress

import (
	"context"
	"fmt"
	neturl "net/url"

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
	baseUrl := fmt.Sprintf("%s.%s", "https://www-radix-canary-golang-prod", cfg.GetClusterFQDN())
	url, _ := neturl.JoinPath(baseUrl, "health")
	client := httpUtils.GetHTTPDefaultClient()
	log.Ctx(ctx).Debug().Str("url", url).Msg("Requesting data")

	// Run tests ingress
	_, err := client.Get(url)
	if err == nil {
		// Successful test
		return nil
	}

	// Failed test
	return err
}

// Success is a function after a call to Reach succeeds
func Success(testName string) {
	nspMetrics.AddIngressReachable()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	logger.Info().Str("test", testName).Msg("Test: SUCCESS")
}

// Fail is a function after a call to Reach failed
func Fail(testName string) {
	nspMetrics.AddIngressUnreachable()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	logger.Info().Str("test", testName).Msg("Test: FAIL")
}
