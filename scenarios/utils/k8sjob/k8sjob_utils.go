package k8sjob

import (
	"fmt"

	jobClient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
)

// IsListedWithStatus Checks if job exists with status
func IsListedWithStatus(cfg config.Config, appName string, appEnv string, jobComponentName string, batchName string, expectedStatus string) error {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()
	params := jobClient.NewGetBatchesParams().
		WithJobComponentName(jobComponentName).
		WithAppName(appName).
		WithEnvName(appEnv).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetK8sJobClient(cfg)

	batches, err := client.GetBatches(params, clientBearerToken)

	if err != nil {
		return fmt.Errorf("error calling GetBatches for application %s in environment %s: %v", appName, appEnv, err)
	}

	for _, batchSummary := range batches.Payload {
		if *batchSummary.Name == batchName && *batchSummary.Status == expectedStatus {
			return nil
		}
	}

	return fmt.Errorf("could not find batch job %s with expected status %s", batchName, expectedStatus)
}
