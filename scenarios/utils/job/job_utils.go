package job

import (
	"fmt"

	pipelineJobClient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/pipeline_job"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

// GetLastPipelineJobWithStatus Checks if a last pipeline job exists with status
func GetLastPipelineJobWithStatus(cfg config.Config, appName, expectedStatus string, logger *log.Entry) (*models.JobSummary, error) {
	jobSummaries, err := getPipelineJobs(cfg, appName)
	if err != nil {
		return nil, err
	}
	lastJobSummary := jobSummaries[0]
	if lastJobSummary.Status != expectedStatus {
		return nil, fmt.Errorf("method GetLastPipelineJobWithStatus for application %s expected status '%s', but it received '%s'",
			appName, expectedStatus, lastJobSummary.Status)
	}
	logger.Debugf("method GetLastPipelineJobWithStatus for application %s received expected status '%s'", appName, expectedStatus)
	return lastJobSummary, nil
}

// GetAnyPipelineJobWithStatus Checks if any pipeline job exists with status
func GetAnyPipelineJobWithStatus(cfg config.Config, appName, expectedStatus string, logger *log.Entry) (*models.JobSummary, error) {
	jobSummaries, err := getPipelineJobs(cfg, appName)
	if err != nil {
		return nil, err
	}
	for _, jobSummary := range jobSummaries {
		if jobSummary.Status == expectedStatus {
			logger.Debugf("method GetAnyPipelineJobWithStatus for application %s received expected status '%s'", appName, expectedStatus)
			return jobSummary, nil
		}
	}
	return nil, fmt.Errorf("method GetAnyPipelineJobWithStatus for application %s expected any job with the status '%s', but it does not exist",
		appName, expectedStatus)
}

func getPipelineJobs(cfg config.Config, appName string) ([]*models.JobSummary, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := pipelineJobClient.NewGetApplicationJobsParams().
		WithAppName(appName).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetJobClient(cfg)

	applicationJobs, err := client.GetApplicationJobs(params, clientBearerToken)
	if err != nil {
		return nil, fmt.Errorf("error calling GetApplicationJobs for application %s: %v", appName, err)
	}
	if applicationJobs.Payload == nil || len(applicationJobs.Payload) == 0 {
		return nil, fmt.Errorf("method GetApplicationJobs for application %s received invalid or empty applicationJobs payload", appName)
	}
	jobSummaries := applicationJobs.Payload
	return jobSummaries, nil
}

// Stop Stops a job
func Stop(cfg config.Config, appName, jobName string) error {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetJobClient(cfg)

	params := pipelineJobClient.NewStopApplicationJobParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	jobStopped, err := client.StopApplicationJob(params, clientBearerToken)
	if err == nil && jobStopped != nil {
		return nil
	}

	return fmt.Errorf("failed to stop job %s for an app %s. Error: %w", jobName, appName, err)
}

// IsDone Checks if job is done
func IsDone(cfg config.Config, appName, jobName string, logger *log.Entry) (string, error) {
	jobStatus, err := GetStatus(cfg, appName, jobName, logger)
	if err != nil {
		return "", err
	}
	if jobStatus == "Succeeded" || jobStatus == "Failed" {
		logger.Debugf("Job %s for an app %s is done with status: %s", jobName, appName, jobStatus)
		return jobStatus, nil
	}
	logger.Debugf("Job %s for an app %s is not done yet", jobName, appName)
	return "", fmt.Errorf("job %s for an app %s is not complete yet, Status %s", jobName, appName, jobStatus)
}

// GetStatus Gets status of job
func GetStatus(cfg config.Config, appName, jobName string, logger *log.Entry) (string, error) {
	job, err := Get(cfg, appName, jobName)
	if err != nil {
		return "", err
	}
	if job != nil {
		return job.Status, nil
	}
	logger.Debugf("Job %s was not listed yet", jobName)
	return "", fmt.Errorf("job %s does not exist", jobName)
}

// Get gets job from job name
func Get(cfg config.Config, appName, jobName string) (*models.Job, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := pipelineJobClient.NewGetApplicationJobParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetJobClient(cfg)

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
func GetSteps(cfg config.Config, appName, jobName string) []*models.Step {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := pipelineJobClient.NewGetApplicationJobParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetJobClient(cfg)

	applicationJob, err := client.GetApplicationJob(params, clientBearerToken)
	if err == nil &&
		applicationJob.Payload != nil &&
		applicationJob.Payload.Steps != nil {

		return applicationJob.Payload.Steps
	}

	return nil
}

// GetLogForStep gets log for step
func GetLogForStep(cfg config.Config, appName, jobName, stepName string, logger *log.Entry) string {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := pipelineJobClient.NewGetPipelineJobStepLogsParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithStepName(stepName).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetJobClient(cfg)

	applicationJobLogs, err := client.GetPipelineJobStepLogs(params, clientBearerToken)
	if err != nil {
		logger.Errorf("failed to get pipeline log for the app %s, job %s, step %s", appName, jobName, stepName)
		return ""
	}
	return applicationJobLogs.Payload
}
