package egresspolicy

import (
	"context"
	"fmt"

	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// LookupInternalDNS tests that we are able to make lookups to internal DNS
func LookupInternalDNS(ctx context.Context, cfg config.Config) error {
	internalDnsUrl := fmt.Sprintf("%s/testinternaldns", cfg.GetNetworkPolicyCanaryUrl("egressrulestopublicdns"))
	return lookupDns(internalDnsUrl, ctx)
}

// LookupPublicDNS tests that we are able to make lookups to public DNS
func LookupPublicDNS(ctx context.Context, cfg config.Config) error {
	publicDnsUrl := fmt.Sprintf("%s/testpublicdns", cfg.GetNetworkPolicyCanaryUrl("egressrulestopublicdns"))
	return lookupDns(publicDnsUrl, ctx)
}

func lookupDns(dnsUrl string, ctx context.Context) error {
	client := httpUtils.GetHTTPDefaultClient()

	log.Ctx(ctx).Debug().Str("url", dnsUrl).Msg("Requesting data")
	dnsResponse, dnsErr := client.Get(dnsUrl)

	if dnsErr != nil {
		return dnsErr
	}
	if dnsResponse.StatusCode != 200 {
		return errors.Errorf("expected dnsResponse.StatusCode is 200, but got %d", dnsResponse.StatusCode)
	}
	return nil
}

// InternalDnsSuccess is a function after a call to Lookup succeeds
func InternalDnsSuccess(ctx context.Context, testName string) {
	nspMetrics.AddInternalDnsIsHealthy()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: SUCCESS")
}

// InternalDnsFail is a function after a call to Lookup failed
func InternalDnsFail(ctx context.Context, testName string) {
	nspMetrics.AddInternalDnsIsUnhealthy()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: FAIL")
}

// PublicDnsSuccess is a function after a call to Lookup succeeds
func PublicDnsSuccess(ctx context.Context, testName string) {
	nspMetrics.AddPublicDnsIsHealthy()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: SUCCESS")
}

// PublicDnsFail is a function after a call to Lookup failed
func PublicDnsFail(ctx context.Context, testName string) {
	nspMetrics.AddPublicDnsIsUnhealthy()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: FAIL")
}
