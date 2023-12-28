package egresspolicy

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/rs/zerolog/log"
)

const (
	timeout = 10
)

// ReachRadixSite tests that canary golang endpoint can be reached from networkpolicy canary with policy that allows it
func ReachRadixSite(ctx context.Context, cfg config.Config) error {
	appEnv := "allowradix"
	reachRadixSiteUrl := getReachRadixSiteUrl(cfg, appEnv)
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	_, err := client.Get(reachRadixSiteUrl)
	if err != nil {
		return err
	}
	return nil
}

// NotReachRadixSite tests that canary golang endpoint can not be reached from networkpolicy canary with policy that prohibits it
func NotReachRadixSite(ctx context.Context, cfg config.Config) error {
	appEnv := "egressrulestopublicdns"
	reachRadixSiteUrl := getReachRadixSiteUrl(cfg, appEnv)
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	res, err := client.Get(reachRadixSiteUrl)
	if err == nil && res.StatusCode == 200 {
		return fmt.Errorf("request to %s from canary should have been blocked by network policy", reachRadixSiteUrl)
	}
	return nil
}

// NotReachExternalSite tests that a list of external websites can not be reached from networkpolicy canary with policy that prohibits it
func NotReachExternalSite(ctx context.Context, cfg config.Config) error {
	appEnvs := []string{"egressrulestopublicdns", "allowradix"}
	var errs []error
	for _, appEnv := range appEnvs {
		reachExternalSiteUrl := fmt.Sprintf("%s/testexternalwebsite", cfg.GetNetworkPolicyCanaryUrl(appEnv))
		client := http.Client{
			Timeout: 60 * time.Second,
		}
		res, err := client.Get(reachExternalSiteUrl)
		if err == nil && res.StatusCode == 200 {
			errs = append(errs, fmt.Errorf("requests to external websites from canary in environment %s should have been blocked by network policy", appEnv))
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func getReachRadixSiteUrl(cfg config.Config, appEnv string) string {
	return fmt.Sprintf("%s/testradixsite", cfg.GetNetworkPolicyCanaryUrl(appEnv))
}

// ReachRadixSiteSuccess is a function after a call to ReachRadixSite succeeds
func ReachRadixSiteSuccess(ctx context.Context, testName string) {
	nspMetrics.AddRadixSiteReachable()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: SUCCESS")
}

// ReachRadixSiteFail is a function after a call to ReachRadixSite failed
func ReachRadixSiteFail(ctx context.Context, testName string) {
	nspMetrics.AddRadixSiteUnreachable()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: FAIL")
}

// NotReachRadixSiteSuccess is a function after a call to NotReachRadixSite succeeds
func NotReachRadixSiteSuccess(ctx context.Context, testName string) {
	nspMetrics.AddNotRadixSiteReachable()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: SUCCESS")
}

// NotReachRadixSiteFail is a function after a call to NotReachRadixSite failed
func NotReachRadixSiteFail(ctx context.Context, testName string) {
	nspMetrics.AddNotRadixSiteUnreachable()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: FAIL")
}

// NotReachExternalSiteSuccess is a function after a call to NotReachExternalSite failed
func NotReachExternalSiteSuccess(ctx context.Context, testName string) {
	nspMetrics.AddNotExternalSiteReachable()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: SUCCESS")
}

// NotReachExternalSiteFail is a function after a call to NotReachExternalSite failed
func NotReachExternalSiteFail(ctx context.Context, testName string) {
	nspMetrics.AddNotExternalSiteUnreachable()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: FAIL")
}
