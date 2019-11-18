package buildsecrets

import (
	jobclient "github.com/equinor/radix-cicd-canary/generated-client/client/job"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/build"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Set Tests that we are able to successfully set build secrets
func Set(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Trigger build to apply RA with build secrets
	ok := httpUtils.TriggerWebhookPush(env, config.App2BranchToBuildFrom, config.App2CommitID, config.App2SSHRepository, config.App2SharedSecret)
	if !ok {
		return false, nil
	}

	logger.Infof("Job was triggered to apply RA")

	// Get job
	ok, _ = test.WaitForCheckFunc(env, jobIsListedAndFailed)
	if !ok {
		return false, nil
	}

	// First job failed, due to missing build secrets, as expected in test
	// Set build secrets

	return true, nil
}

func jobIsListedAndFailed(env env.Env, args []string) (bool, interface{}) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := jobclient.NewGetApplicationJobsParams().
		WithAppName(config.App2Name).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetJobClient(env)

	applicationJobs, err := client.GetApplicationJobs(params, clientBearerToken)
	if err == nil && applicationJobs.Payload != nil && len(applicationJobs.Payload) > 0 {
		jobName := applicationJobs.Payload[0].Name
		job := build.Job(env, jobName)

		// Job has failed status and has no build step
		if job != nil && job.Status == "Failed" && len(job.Steps) == 2 {
			return true, applicationJobs.Payload[0]
		}
	}

	logger.Info("Job was not listed yet")
	return false, nil

}
