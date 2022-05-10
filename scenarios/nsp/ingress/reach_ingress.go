package ingress

import (
	"fmt"
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
	url := fmt.Sprintf("%s/health", env.GetGolangCanaryUrl())
	logger.Debugf("Requesting data from %s", url)

	// Run tests ingress
	_, err := client.Get(url)
	if err == nil {
		// Successful test
		return true, nil
	}

	// Failed test
	return false, err
}

// Success is a function after a call to Reach succeeds
func Success(testName string) {
	nspMetrics.AddIngressReachable()
	metrics.AddTestSuccess(testName, nspMetrics.Success)
	metrics.AddTestNoError(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// Fail is a function after a call to Reach failed
func Fail(testName string) {
	nspMetrics.AddIngressUnreachable()
	metrics.AddTestNoSuccess(testName, nspMetrics.Success)
	metrics.AddTestError(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}
