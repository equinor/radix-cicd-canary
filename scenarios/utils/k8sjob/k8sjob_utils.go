package k8sjob

import (
	jobClient "github.com/equinor/radix-cicd-canary/generated-client/client/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

// IsListedWithStatus Checks if job exists with status
func IsListedWithStatus(env env.Env, appName string, appEnv string, jobComponentName string, batchName string, expectedStatus string) (bool, interface{}) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()
	params := jobClient.NewGetBatchesParams().
		WithJobComponentName(jobComponentName).
		WithAppName(appName).
		WithEnvName(appEnv).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetK8sJobClient(env)

	batches, err := client.GetBatches(params, clientBearerToken)

	if err != nil {
		log.Errorf("Error calling GetBatches for application %s in environment %s: %v", appName, appEnv, err)
		return false, nil
	}

	for _, batchSummary := range batches.Payload {
		if *batchSummary.Name == batchName && *batchSummary.Status == expectedStatus {
			return true, nil
		}
	}

	return false, nil
}
