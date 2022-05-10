package egresspolicy

import (
	"fmt"
	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-common/utils/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	timeout = 10
)

// ReachRadixSite tests that canary golang endpoint can be reached from networkpolicy canary with policy that allows it
func ReachRadixSite(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})
	appEnv := "allowradix"
	reachRadixSiteUrl := getReachRadixSiteUrl(env, appEnv)
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	_, err := client.Get(reachRadixSiteUrl)
	if err != nil {
		return false, err
	}
	return true, nil
}

// NotReachRadixSite tests that canary golang endpoint can not be reached from networkpolicy canary with policy that prohibits it
func NotReachRadixSite(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})
	appEnv := "egressrulestopublicdns"
	reachRadixSiteUrl := getReachRadixSiteUrl(env, appEnv)
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	res, err := client.Get(reachRadixSiteUrl)
	if err == nil && res.StatusCode == 200 {
		return false, fmt.Errorf("request to %s from canary should have been blocked by network policy", env.GetGolangCanaryUrl())
	}
	return true, nil
}

// NotReachExternalSite tests that a list of external websites can not be reached from networkpolicy canary with policy that prohibits it
func NotReachExternalSite(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})
	appEnvs := []string{"egressrulestopublicdns", "allowradix"}
	var errs []error
	for _, appEnv := range appEnvs {
		reachExternalSiteUrl := fmt.Sprintf("%s/testexternalwebsite", env.GetNetworkPolicyCanaryUrl(appEnv))
		client := http.Client{
			Timeout: 60 * time.Second,
		}
		res, err := client.Get(reachExternalSiteUrl)
		if err == nil && res.StatusCode == 200 {
			errs = append(errs, fmt.Errorf("requests to external websites from canary in environment %s should have been blocked by network policy", appEnv))
		}
	}
	if len(errs) > 0 {
		return false, errors.Concat(errs)
	}
	return true, nil
}

func getReachRadixSiteUrl(env env.Env, appEnv string) string {
	return fmt.Sprintf("%s/testradixsite", env.GetNetworkPolicyCanaryUrl(appEnv))
}

// ReachRadixSiteSuccess is a function after a call to ReachRadixSite succeeds
func ReachRadixSiteSuccess(testName string) {
	nspMetrics.AddRadixSiteReachable()
	metrics.AddTestSuccess(testName, nspMetrics.Success)
	metrics.AddTestNoError(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// ReachRadixSiteFail is a function after a call to ReachRadixSite failed
func ReachRadixSiteFail(testName string) {
	nspMetrics.AddRadixSiteUnreachable()
	metrics.AddTestNoSuccess(testName, nspMetrics.Success)
	metrics.AddTestError(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}

// NotReachRadixSiteSuccess is a function after a call to NotReachRadixSite succeeds
func NotReachRadixSiteSuccess(testName string) {
	nspMetrics.AddNotRadixSiteReachable()
	metrics.AddTestSuccess(testName, nspMetrics.Success)
	metrics.AddTestNoError(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// NotReachRadixSiteFail is a function after a call to NotReachRadixSite failed
func NotReachRadixSiteFail(testName string) {
	nspMetrics.AddNotRadixSiteUnreachable()
	metrics.AddTestNoSuccess(testName, nspMetrics.Success)
	metrics.AddTestError(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}

// NotReachExternalSiteSuccess is a function after a call to NotReachExternalSite failed
func NotReachExternalSiteSuccess(testName string) {
	nspMetrics.AddNotExternalSiteReachable()
	metrics.AddTestSuccess(testName, nspMetrics.Success)
	metrics.AddTestNoError(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// NotReachExternalSiteFail is a function after a call to NotReachExternalSite failed
func NotReachExternalSiteFail(testName string) {
	nspMetrics.AddNotExternalSiteUnreachable()
	metrics.AddTestNoSuccess(testName, nspMetrics.Success)
	metrics.AddTestError(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}
