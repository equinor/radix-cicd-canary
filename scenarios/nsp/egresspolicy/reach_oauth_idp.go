package egresspolicy

import (
	"fmt"
	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// ReachOauthIdp tests that IDP endpoint can be reached from Oauth Aux pod
func ReachOauthIdp(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})
	appEnv := "oauthdenyall"
	timeout := 15
	oauthCallbackUrl := fmt.Sprintf("%s/oauth2/callback?code=bullshitcode", env.GetNetworkPolicyCanaryUrl(appEnv))
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	_, err := client.Get(oauthCallbackUrl)
	if err == http.ErrHandlerTimeout {
		return false, fmt.Errorf("got no response from /oauth/callback within %d seconds, which likely means oauth pod could not connect to IDP. should be allowed by nsp", timeout)
	}
	return true, nil
}

// ReachOauthIdpSuccess is a function after a call to ReachOauthIdp succeeds
func ReachOauthIdpSuccess(testName string) {
	nspMetrics.AddOauthIdpReachable()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// ReachOauthIdpFail is a function after a call to ReachOauthIdp failed
func ReachOauthIdpFail(testName string) {
	nspMetrics.AddOauthIdpUnreachable()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}