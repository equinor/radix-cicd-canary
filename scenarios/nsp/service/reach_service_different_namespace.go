package service

import (
	"errors"

	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Reach tests that we are able to reach radix-canary-golang-prod endpoint
func Reach(cfg config.Config, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	client := httpUtils.GetHTTPDefaultClient()
	url := "http://www.radix-canary-golang-prod:5000/health"
	logger.Infof("requesting data from %s", url)

	// Run tests service
	_, err := client.Get(url)
	if err == nil {
		// Failing test
		return errors.New("no error was returned when attempting to access service")
	}

	// Successful
	return nil
}

// Success is a function after a call to Reach succeeds
func Success(testName string) {
	nspMetrics.AddServiceUnreachable()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// Fail is a function after a call to Reach failed
func Fail(testName string) {
	nspMetrics.AddServiceReachable()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}
