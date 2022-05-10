package egresspolicy

import (
	"fmt"
	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// LookupInternalDNS tests that we are able to make lookups to internal DNS
func LookupInternalDNS(env env.Env, suiteName string) (bool, error) {
	internalDnsUrl := fmt.Sprintf("%s/testinternaldns", env.GetNetworkPolicyCanaryUrl("egressrulestopublicdns"))
	return lookupDns(internalDnsUrl, suiteName)
}

// LookupPublicDNS tests that we are able to make lookups to public DNS
func LookupPublicDNS(env env.Env, suiteName string) (bool, error) {
	publicDnsUrl := fmt.Sprintf("%s/testpublicdns", env.GetNetworkPolicyCanaryUrl("egressrulestopublicdns"))
	return lookupDns(publicDnsUrl, suiteName)
}

func lookupDns(dnsUrl string, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	client := httpUtils.GetHTTPDefaultClient()

	logger.Debugf("Requesting data from %s", dnsUrl)
	dnsResponse, dnsErr := client.Get(dnsUrl)

	if dnsErr != nil {
		return false, dnsErr
	}
	return dnsResponse.StatusCode == 200, nil
}

// InternalDnsSuccess is a function after a call to Lookup succeeds
func InternalDnsSuccess(testName string) {
	nspMetrics.AddInternalDnsIsHealthy()
	metrics.AddTestSuccess(testName, nspMetrics.Success)
	metrics.AddTestNoError(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// InternalDnsFail is a function after a call to Lookup failed
func InternalDnsFail(testName string) {
	nspMetrics.AddInternalDnsIsUnhealthy()
	metrics.AddTestNoSuccess(testName, nspMetrics.Success)
	metrics.AddTestError(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}

// PublicDnsSuccess is a function after a call to Lookup succeeds
func PublicDnsSuccess(testName string) {
	nspMetrics.AddPublicDnsIsHealthy()
	metrics.AddTestSuccess(testName, nspMetrics.Success)
	metrics.AddTestNoError(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// PublicDnsFail is a function after a call to Lookup failed
func PublicDnsFail(testName string) {
	nspMetrics.AddPublicDnsIsUnhealthy()
	metrics.AddTestNoSuccess(testName, nspMetrics.Success)
	metrics.AddTestError(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}
