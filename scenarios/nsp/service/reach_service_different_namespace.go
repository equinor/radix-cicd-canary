package service

import (
	"errors"
	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Reach tests that we are able to reach radix-canary-golang-prod endpoint
func Reach(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	client := httpUtils.GetHTTPDefaultClient()
	url := "http://www.radix-canary-golang-prod:5000/health"
	logger.Infof("requesting data from %s", url)

	// Run tests service
	_, err := client.Get(url)
	if err == nil {
		// Failing test
		return false, errors.New("No error was returned when attempting to access service")
	}

	// Successful
	return true, nil
}

// Success is a function after a call to Reach succeeds
func Success(testName string) {
	nspMetrics.AddServiceUnreachable()
	metrics.AddTestSuccess(testName, nspMetrics.Success)
	metrics.AddTestNoError(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// Fail is a function after a call to Reach failed
func Fail(testName string) {
	nspMetrics.AddServiceReachable()
	metrics.AddTestNoSuccess(testName, nspMetrics.Success)
	metrics.AddTestError(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}
