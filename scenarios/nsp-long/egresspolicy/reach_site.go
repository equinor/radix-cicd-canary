package egresspolicy

import (
	"fmt"
	"net/http"
	"time"

	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-common/utils/errors"
	log "github.com/sirupsen/logrus"
)

const (
	timeout = 10
)

// ReachRadixSite tests that canary golang endpoint can be reached from networkpolicy canary with policy that allows it
func ReachRadixSite(cfg config.Config, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})
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
func NotReachRadixSite(cfg config.Config, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})
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
func NotReachExternalSite(cfg config.Config, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})
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
		return errors.Concat(errs)
	}
	return nil
}

func getReachRadixSiteUrl(cfg config.Config, appEnv string) string {
	return fmt.Sprintf("%s/testradixsite", cfg.GetNetworkPolicyCanaryUrl(appEnv))
}

// ReachRadixSiteSuccess is a function after a call to ReachRadixSite succeeds
func ReachRadixSiteSuccess(testName string) {
	nspMetrics.AddRadixSiteReachable()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// ReachRadixSiteFail is a function after a call to ReachRadixSite failed
func ReachRadixSiteFail(testName string) {
	nspMetrics.AddRadixSiteUnreachable()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}

// NotReachRadixSiteSuccess is a function after a call to NotReachRadixSite succeeds
func NotReachRadixSiteSuccess(testName string) {
	nspMetrics.AddNotRadixSiteReachable()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// NotReachRadixSiteFail is a function after a call to NotReachRadixSite failed
func NotReachRadixSiteFail(testName string) {
	nspMetrics.AddNotRadixSiteUnreachable()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}

// NotReachExternalSiteSuccess is a function after a call to NotReachExternalSite failed
func NotReachExternalSiteSuccess(testName string) {
	nspMetrics.AddNotExternalSiteReachable()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// NotReachExternalSiteFail is a function after a call to NotReachExternalSite failed
func NotReachExternalSiteFail(testName string) {
	nspMetrics.AddNotExternalSiteUnreachable()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}
