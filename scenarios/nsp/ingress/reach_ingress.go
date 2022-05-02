package ingress

import (
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Reach tests that we are able to reach radix-canary-golang-prod endpoint
func Reach(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	client := httpUtils.GetHTTPDefaultClient()
	url := fmt.Sprintf("%s/health", getIngressForRadixCanaryApp(env.GetClusterFQDN()))
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

func getIngressForRadixCanaryApp(clusterFQDN string) string {
	canaryURLPrefix := "https://www-radix-canary-golang-prod"
	return fmt.Sprintf("%s.%s", canaryURLPrefix, clusterFQDN)
}
