package build

import (
	jobclient "github.com/equinor/radix-cicd-canary/generated-client/client/job"
	models "github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

// Application Tests that we are able to successfully build an application
func Application(env env.Env) (bool, error) {
	// Trigger build via web hook
	ok := httpUtils.TriggerWebhookPush(env, config.App2BranchToBuildFrom, config.App2CommitID, config.App2SSHRepository, config.App2SharedSecret)
	if !ok {
		return false, nil
	}

	// Get job
	ok, jobSummary := test.WaitForCheckFunc(env, isJobListed)
	if ok {
		jobName := (jobSummary.(*models.JobSummary)).Name

		// Another build should cause second job to queue up
		// Trigger another build via web hook
		ok := httpUtils.TriggerWebhookPush(env, config.App2BranchToBuildFrom, config.App2CommitID, config.App2SSHRepository, config.App2SharedSecret)
		if !ok {
			return false, nil
		}

		ok, _ = test.WaitForCheckFuncWithArguments(env, isSecondJobExpectedStatus, []string{"Queued"})
		if !ok {
			return false, nil
		}

		log.Info("Second job was queued")
		ok, status := test.WaitForCheckFuncWithArguments(env, isJobDone, []string{jobName})

		if ok && status.(string) == "Succeeded" {
			ok, _ = test.WaitForCheckFuncWithArguments(env, isSecondJobExpectedStatus, []string{"Running"})
			if !ok {
				return false, nil
			}

			return true, nil
		}

	}

	return false, nil
}

func isJobListed(env env.Env, args []string) (bool, interface{}) {
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
		return true, applicationJobs.Payload[0]
	}

	log.Info("Job was not listed yet")
	return false, nil
}

func isSecondJobExpectedStatus(env env.Env, args []string) (bool, interface{}) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	expectedStatus := args[0]
	params := jobclient.NewGetApplicationJobsParams().
		WithAppName(config.App2Name).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetJobClient(env)

	applicationJobs, err := client.GetApplicationJobs(params, clientBearerToken)
	if err == nil && applicationJobs.Payload != nil && len(applicationJobs.Payload) > 0 && applicationJobs.Payload[0].Status == expectedStatus {
		return true, nil
	}

	log.Info("Queued job was not listed yet")
	return false, nil
}

func isJobDone(env env.Env, args []string) (bool, interface{}) {
	jobStatus := getJobStatus(env, args[0])
	if jobStatus == "Succeeded" || jobStatus == "Failed" {
		log.Info("Job is done")
		return true, jobStatus
	}

	log.Info("Job is not done yet")
	return false, nil
}

func getJobStatus(env env.Env, jobName string) string {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := jobclient.NewGetApplicationJobParams().
		WithAppName(config.App2Name).
		WithJobName(jobName).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetJobClient(env)

	applicationJob, err := client.GetApplicationJob(params, clientBearerToken)
	if err == nil && applicationJob.Payload != nil {
		return applicationJob.Payload.Status
	}

	log.Info("Job was not listed yet")
	return ""
}
