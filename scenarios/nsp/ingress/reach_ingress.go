package ingress

import (
	"context"
	"fmt"
	neturl "net/url"

	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/rs/zerolog/log"
)

// Reach tests that we are able to reach radix-canary-golang-prod endpoint
func Reach(ctx context.Context, cfg config.Config) error {
	baseUrl := fmt.Sprintf("%s.%s", "https://www-radix-canary-golang-prod", cfg.GetClusterFQDN())
	url, _ := neturl.JoinPath(baseUrl, "health")
	client := httpUtils.GetHTTPDefaultClient(cfg.GetNSPReachIngressTimeout())
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
func Success(ctx context.Context, testName string) {
	nspMetrics.AddIngressReachable()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: SUCCESS")
}

// Fail is a function after a call to Reach failed
func Fail(ctx context.Context, testName string) {
	nspMetrics.AddIngressUnreachable()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: FAIL")
}
