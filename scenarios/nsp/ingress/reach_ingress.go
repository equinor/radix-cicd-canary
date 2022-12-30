package ingress

import (
	"fmt"
	neturl "net/url"

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

	baseUrl := fmt.Sprintf("%s.%s", "https://www-radix-canary-golang-prod", cfg.GetClusterFQDN())
	url, _ := neturl.JoinPath(baseUrl, "health")
	client := httpUtils.GetHTTPDefaultClient()
	logger.Debugf("Requesting data from %s", url)

	// Run tests ingress
	_, err := client.Get(url)
	if err == nil {
		// Successful test
		return nil
	}

	// Failed test
	return err
}

// Success is a function after a call to Reach succeeds
func Success(testName string) {
	nspMetrics.AddIngressReachable()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// Fail is a function after a call to Reach failed
func Fail(testName string) {
	nspMetrics.AddIngressUnreachable()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}
