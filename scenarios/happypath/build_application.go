package happypath

import (
	applicationclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	jobclient "github.com/equinor/radix-cicd-canary/generated-client/client/job"
	models "github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

func buildApplication(env env.Env) (bool, error) {
	test.WaitForCheckFunc(env, isApplicationDefined)

	// Trigger build via web hook
	ok := httpUtils.TriggerWebhookPush(env, app2BranchToBuildFrom, app2CommitID, app2SSHRepository, app2SharedSecret)
	if !ok {
		return false, nil
	}

	// Get job
	ok, jobSummary := test.WaitForCheckFunc(env, isJobListed)
	if ok {
		jobName := (jobSummary.(*models.JobSummary)).Name
		ok, status := test.WaitForCheckFuncWithArguments(env, isJobDone, []string{jobName})

		if ok && status.(string) == "Succeeded" {
			return true, nil
		}

	}

	return false, nil
}

func isApplicationDefined(env env.Env, args []string) (bool, interface{}) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := applicationclient.NewGetApplicationParams().
		WithAppName(app2Name).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	_, err := client.GetApplication(params, clientBearerToken)
	if err == nil {
		return true, nil
	}

	log.Info("Application is not defined")
	return false, nil
}

func isJobListed(env env.Env, args []string) (bool, interface{}) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := jobclient.NewGetApplicationJobsParams().
		WithAppName(app2Name).
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
		WithAppName(app2Name).
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
