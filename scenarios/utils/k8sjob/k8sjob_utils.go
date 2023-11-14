package k8sjob

import (
	jobClient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/pkg/errors"
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

	client := httpUtils.GetK8sJobClient(cfg)
	batches, err := client.GetBatches(params, nil)

	if err != nil {
		return errors.Wrapf(err, "error calling GetBatches for application %s in environment %s", appName, appEnv)
	}

	for _, batchSummary := range batches.Payload {
		if *batchSummary.Name == batchName && *batchSummary.Status == expectedStatus {
			return nil
		}
	}

	return errors.Errorf("could not find batch job %s with expected status %s", batchName, expectedStatus)
}
