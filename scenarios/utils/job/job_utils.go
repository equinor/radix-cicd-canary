package job

import (
	"fmt"

	pipelineJobClient "github.com/equinor/radix-cicd-canary/generated-client/client/pipeline_job"
	"github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

// IsListedWithStatus Checks if job exists with status
func IsListedWithStatus(env env.Env, appName, expectedStatus string) (*models.JobSummary, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := pipelineJobClient.NewGetApplicationJobsParams().
		WithAppName(appName).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetJobClient(env)

	applicationJobs, err := client.GetApplicationJobs(params, clientBearerToken)
	if err != nil {
		return nil, fmt.Errorf("error calling GetApplicationJobs for application %s: %v", appName, err)
	}
	if applicationJobs.Payload == nil || len(applicationJobs.Payload) == 0 {
		return nil, fmt.Errorf("method GetApplicationJobs for application %s received invalid or empty applicationJobs payload", appName)
	}
	if applicationJobs.Payload[0].Status != expectedStatus {
		return nil, fmt.Errorf("method GetApplicationJobs for application %s expected status '%s', but it received '%s'",
			appName, expectedStatus, applicationJobs.Payload[0].Status)
	}
	log.Debugf("method GetApplicationJobs for application %s received expected status '%s'", appName, expectedStatus)
	return applicationJobs.Payload[0], nil
}

// Stop Stops a job
func Stop(env env.Env, appName, jobName string) error {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetJobClient(env)

	params := pipelineJobClient.NewStopApplicationJobParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	jobStopped, err := client.StopApplicationJob(params, clientBearerToken)
	if err == nil && jobStopped != nil {
		return nil
	}

	return fmt.Errorf("failed to stop job %s. Error: %w", jobName, err)
}

// IsDone Checks if job is done
func IsDone(appName, jobName string, env env.Env) (string, error) {
	jobStatus, err := GetStatus(env, appName, jobName)
	if err != nil {
		return "", err
	}
	if jobStatus == "Succeeded" || jobStatus == "Failed" {
		log.Debugf("Job is done with status: %s", jobStatus)
		return jobStatus, nil
	}
	log.Debug("Job is not done yet")
	return "", fmt.Errorf("job %s was possible failed, Status %s", jobName, jobStatus)
}

// GetStatus Gets status of job
func GetStatus(env env.Env, appName, jobName string) (string, error) {
	job, err := Get(env, appName, jobName)
	if err != nil {
		return "", err
	}
	if job != nil {
		return job.Status, nil
	}
	log.Debug("Job was not listed yet")
	return "", fmt.Errorf("job %s does not exist", jobName)
}

// Get gets job from job name
func Get(env env.Env, appName, jobName string) (*models.Job, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := pipelineJobClient.NewGetApplicationJobParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetJobClient(env)

	applicationJob, err := client.GetApplicationJob(params, clientBearerToken)
	if err != nil {
		return nil, err
	}
	if applicationJob.Payload != nil {
		return applicationJob.Payload, nil
	}
	return nil, fmt.Errorf("failed to get job %s", jobName)
}

// GetSteps gets job from job name
func GetSteps(env env.Env, appName, jobName string) []*models.Step {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := pipelineJobClient.NewGetApplicationJobParams().
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

	params := pipelineJobClient.NewGetPipelineJobStepLogsParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithStepName(stepName).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetJobClient(env)

	applicationJobLogs, err := client.GetPipelineJobStepLogs(params, clientBearerToken)
	if err != nil {
		log.Errorf("failed to get pipeline log for the app %s, job %s, step %s", appName, jobName, stepName)
		return ""
	}
	return applicationJobLogs.Payload
}
