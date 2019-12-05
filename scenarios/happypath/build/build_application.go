package build

import (
	"strings"
	"time"

	jobclient "github.com/equinor/radix-cicd-canary/generated-client/client/job"
	models "github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/array"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

const (
	Secret1      = "SECRET_1"
	Secret2      = "SECRET_2"
	Secret1Value = "SECRET_1_VALUE"
	Secret2Value = "SECRET_2_VALUE"
)

type expectedStep struct {
	name       string
	components []string
}

// Application Tests that we are able to successfully build an application
func Application(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Trigger build via web hook
	ok := httpUtils.TriggerWebhookPush(env, config.App2BranchToBuildFrom, config.App2CommitID, config.App2SSHRepository, config.App2SharedSecret)
	if !ok {
		return false, nil
	}
	logger.Infof("First job was triggered")

	// Get job
	ok, jobSummary := test.WaitForCheckFuncWithArguments(env, IsJobListedWithStatus, []string{"Running"})
	if !ok {
		return false, nil
	}

	jobName := (jobSummary.(*models.JobSummary)).Name
	logger.Infof("First job name: %s", jobName)

	// Another build should cause second job to queue up
	// Trigger another build via web hook
	time.Sleep(1 * time.Second)
	ok = httpUtils.TriggerWebhookPush(env, config.App2BranchToBuildFrom, config.App2CommitID, config.App2SSHRepository, config.App2SharedSecret)
	if !ok {
		return false, nil
	}
	logger.Infof("Second job was triggered")

	ok, jobSummary = test.WaitForCheckFuncWithArguments(env, IsJobListedWithStatus, []string{"Queued"})
	if !ok {
		return false, nil
	}

	logger.Info("Second job was queued")
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
		{name: "radix-pipeline", components: []string{}},
		{name: "clone", components: []string{}},
		{name: "build-app", components: []string{"app"}},
		{name: "build-redis", components: []string{"redis"}},
		{name: "scan-app", components: []string{"app"}},
		{name: "scan-redis", components: []string{"redis"}}}

	if steps == nil && len(steps) != len(expectedSteps) {
		logger.Error("Pipeline steps was not as expected")
		return false, nil
	}

	for index, step := range steps {
		if !strings.EqualFold(step.Name, expectedSteps[index].name) {
			logger.Errorf("Expeced step %s, but got %s", expectedSteps[index].name, step.Name)
			return false, nil
		}

		if !array.EqualElements(step.Components, expectedSteps[index].components) {
			logger.Errorf("Expeced components %s, but got %s", expectedSteps[index].components, step.Components)
			return false, nil
		}
	}

	log := getJobLogForStep(env, jobName, "build-app")
	if !strings.Contains(log, Secret1Value) || !strings.Contains(log, Secret2Value) {
		logger.Error("Build secrets are not contained in build log")
		return false, nil
	}

	ok, jobSummary = test.WaitForCheckFuncWithArguments(env, IsJobListedWithStatus, []string{"Running"})
	if !ok {
		return false, nil
	}

	// Stop job and verify that it has been stopped
	jobName = (jobSummary.(*models.JobSummary)).Name
	logger.Infof("Second job name: %s", jobName)
	ok = stopJob(env, config.App2Name, jobName)
	if !ok {
		return false, nil
	}

	ok, jobSummary = test.WaitForCheckFuncWithArguments(env, IsJobListedWithStatus, []string{"Stopped"})
	if !ok {
		return false, nil
	}

	logger.Info("Second job was stopped")
	return true, nil
}

func stopJob(env env.Env, appName, jobName string) bool {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetJobClient(env)

	params := jobclient.NewStopApplicationJobParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	jobStopped, err := client.StopApplicationJob(params, clientBearerToken)
	if err == nil && jobStopped != nil {
		return true
	}

	logger.Infof("Failed stopping job %s. Error: %v", jobName, err)
	return false
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
func getJobLogForStep(env env.Env, jobName, stepName string) string {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := jobclient.NewGetApplicationJobLogsParams().
		WithAppName(config.App2Name).
		WithJobName(jobName).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetJobClient(env)

	applicationJobLogs, err := client.GetApplicationJobLogs(params, clientBearerToken)
	if err == nil &&
		applicationJobLogs.Payload != nil {

		for _, stepLog := range applicationJobLogs.Payload {
			if strings.EqualFold(*stepLog.Name, stepName) {
				return stepLog.Log
			}
		}
	}

	return ""
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
