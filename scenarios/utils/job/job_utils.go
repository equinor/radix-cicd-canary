package job

import (
	"strings"

	jobclient "github.com/equinor/radix-cicd-canary/generated-client/client/job"
	"github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

// IsListedWithStatus Checks if job exists with status
func IsListedWithStatus(env env.Env, appName, expectedStatus string) (bool, interface{}) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := jobclient.NewGetApplicationJobsParams().
		WithAppName(appName).
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

// Stop Stops a job
func Stop(env env.Env, appName, jobName string) bool {
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

	log.Infof("Failed stopping job %s. Error: %v", jobName, err)
	return false
}

// IsDone Checks if job is done
func IsDone(env env.Env, appName, jobName string) (bool, interface{}) {
	jobStatus := GetStatus(env, appName, jobName)
	if jobStatus == "Succeeded" || jobStatus == "Failed" {
		log.Debugf("Job is done with status: %s", jobStatus)
		return true, jobStatus
	}

	log.Debug("Job is not done yet")
	return false, nil
}

// GetStatus Gets status of job
func GetStatus(env env.Env, appName, jobName string) string {
	job := Get(env, appName, jobName)
	if job != nil {
		return job.Status
	}

	log.Debug("Job was not listed yet")
	return ""
}

// Get gets job from job name
func Get(env env.Env, appName, jobName string) *models.Job {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := jobclient.NewGetApplicationJobParams().
		WithAppName(appName).
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

// GetSteps gets job from job name
func GetSteps(env env.Env, appName, jobName string) []*models.Step {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := jobclient.NewGetApplicationJobParams().
		WithAppName(appName).
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

// GetLogForStep gets log for step
func GetLogForStep(env env.Env, appName, jobName, stepName string) string {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := jobclient.NewGetApplicationJobLogsParams().
		WithAppName(appName).
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
