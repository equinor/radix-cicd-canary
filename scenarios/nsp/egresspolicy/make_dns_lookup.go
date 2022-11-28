package egresspolicy

import (
	"fmt"

	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// LookupInternalDNS tests that we are able to make lookups to internal DNS
func LookupInternalDNS(cfg config.Config, suiteName string) error {
	internalDnsUrl := fmt.Sprintf("%s/testinternaldns", cfg.GetNetworkPolicyCanaryUrl("egressrulestopublicdns"))
	return lookupDns(internalDnsUrl, suiteName)
}

// LookupPublicDNS tests that we are able to make lookups to public DNS
func LookupPublicDNS(cfg config.Config, suiteName string) error {
	publicDnsUrl := fmt.Sprintf("%s/testpublicdns", cfg.GetNetworkPolicyCanaryUrl("egressrulestopublicdns"))
	return lookupDns(publicDnsUrl, suiteName)
}

func lookupDns(dnsUrl string, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	client := httpUtils.GetHTTPDefaultClient()

	logger.Debugf("Requesting data from %s", dnsUrl)
	dnsResponse, dnsErr := client.Get(dnsUrl)

	if dnsErr != nil {
		return dnsErr
	}
	if dnsResponse.StatusCode != 200 {
		return fmt.Errorf("expected dnsResponse.StatusCode is 200, but got %d", dnsResponse.StatusCode)
	}
	return nil
}

// InternalDnsSuccess is a function after a call to Lookup succeeds
func InternalDnsSuccess(testName string) {
	nspMetrics.AddInternalDnsIsHealthy()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// InternalDnsFail is a function after a call to Lookup failed
func InternalDnsFail(testName string) {
	nspMetrics.AddInternalDnsIsUnhealthy()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}

// PublicDnsSuccess is a function after a call to Lookup succeeds
func PublicDnsSuccess(testName string) {
	nspMetrics.AddPublicDnsIsHealthy()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// PublicDnsFail is a function after a call to Lookup failed
func PublicDnsFail(testName string) {
	nspMetrics.AddPublicDnsIsUnhealthy()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}
