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

var logger *log.Entry

type expectedStep struct {
	name       string
	components []string
}

// Application Tests that we are able to successfully build an application
func Application(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Trigger build via web hook
	ok := httpUtils.TriggerWebhookPush(env, config.App3BranchToBuildFrom, config.App3CommitID, config.App3SSHRepository, config.App3SharedSecret)
	if !ok {
		return false, nil
	}

	// Get job
	ok, jobSummary := test.WaitForCheckFuncWithArguments(env, IsJobListedWithStatus, []string{"Running"})
	if !ok {
		return false, nil
	}

	jobName := (jobSummary.(*models.JobSummary)).Name

	ok, status := test.WaitForCheckFuncWithArguments(env, isJobDone, []string{jobName})
	if !ok {
		return false, nil
	}
	if status.(string) != "Succeeded" {
		return false, nil
	}

	logger.Info("First job was completed")
	steps := getStepsForJob(env, jobName)

	expectedSteps := []expectedStep{
		{name: "clone-config", components: []string{}},
		{name: "radix-pipeline", components: []string{}}}

	if steps == nil && len(steps) != len(expectedSteps) {
		logger.Error("Pipeline steps was not as expected")
		return false, nil
	}

	return true, nil
}

// IsJobListedWithStatus Checks if job exists with status
func IsJobListedWithStatus(env env.Env, args []string) (bool, interface{}) {
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
	if err == nil && applicationJobs.Payload != nil &&
		len(applicationJobs.Payload) > 0 &&
		applicationJobs.Payload[0].Status == expectedStatus {
		return true, applicationJobs.Payload[0]
	}

	return false, nil
}

func isJobDone(env env.Env, args []string) (bool, interface{}) {
	jobStatus := getJobStatus(env, args[0])
	if jobStatus == "Succeeded" || jobStatus == "Failed" {
		logger.Infof("Job is done with status: %s", jobStatus)
		return true, jobStatus
	}

	logger.Info("Job is not done yet")
	return false, nil
}

func getJobStatus(env env.Env, jobName string) string {
	job := Job(env, jobName)
	if job != nil {
		return job.Status
	}

	logger.Info("Job was not listed yet")
	return ""
}

// Job gets job from job name
func Job(env env.Env, jobName string) *models.Job {
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
		return applicationJob.Payload
	}

	return nil
}

// Job gets job from job name
func getStepsForJob(env env.Env, jobName string) []*models.Step {
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
	if err == nil &&
		applicationJob.Payload != nil &&
		applicationJob.Payload.Steps != nil {

		return applicationJob.Payload.Steps
	}

	return nil
}
