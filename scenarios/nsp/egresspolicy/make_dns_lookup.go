package egresspolicy

import (
	"fmt"
	metrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// LookupInternalDNS tests that we are able to make lookups to internal DNS
func LookupInternalDNS(env env.Env, suiteName string) (bool, error) {
	internalDnsUrl := fmt.Sprintf("%s/testinternaldns", getIngressForRadixCanaryApp(env.GetClusterFQDN()))
	return lookupDns(internalDnsUrl, suiteName)
}

// LookupPublicDNS tests that we are able to make lookups to public DNS
func LookupPublicDNS(env env.Env, suiteName string) (bool, error) {
	publicDnsUrl := fmt.Sprintf("%s/testpublicdns", getIngressForRadixCanaryApp(env.GetClusterFQDN()))
	return lookupDns(publicDnsUrl, suiteName)
}

func lookupDns(dnsUrl string, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	client := httpUtils.GetHTTPDefaultClient()

	logger.Debugf("Requesting data from %s", dnsUrl)
	dnsResponse, dnsErr := client.Get(dnsUrl)

	if dnsErr != nil {
		return false, dnsErr
	} else {
		return dnsResponse.StatusCode == 200, dnsErr
	}
}

// InternalDnsSuccess is a function after a call to Lookup succeeds
func InternalDnsSuccess(testName string) {
	metrics.AddInternalDnsIsHealthy()
	logger.Infof("Test %s: SUCCESS", testName)
}

// InternalDnsFail is a function after a call to Lookup failed
func InternalDnsFail(testName string) {
	metrics.AddInternalDnsIsUnhealthy()
	logger.Infof("Test %s: FAIL", testName)
}

// PublicDnsSuccess is a function after a call to Lookup succeeds
func PublicDnsSuccess(testName string) {
	metrics.AddPublicDnsIsHealthy()
	logger.Infof("Test %s: SUCCESS", testName)
}

// PublicDnsFail is a function after a call to Lookup failed
func PublicDnsFail(testName string) {
	metrics.AddPublicDnsIsUnhealthy()
	logger.Infof("Test %s: FAIL", testName)
}

func getIngressForRadixCanaryApp(clusterFQDN string) string {
	canaryURLPrefix := "https://web-radix-networkpolicy-canary-egressrulestopublicdns"
	return fmt.Sprintf("%s.%s", canaryURLPrefix, clusterFQDN)
}
