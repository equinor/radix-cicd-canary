package job

import (
	"fmt"

	pipelineJobClient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/pipeline_job"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/rs/zerolog"
)

// GetLastPipelineJobWithStatus Checks if a last pipeline job exists with status
func GetLastPipelineJobWithStatus(cfg config.Config, appName, expectedStatus string, logger zerolog.Logger) (*models.JobSummary, error) {
	jobSummaries, err := getPipelineJobs(cfg, appName)
	if err != nil {
		return nil, err
	}
	lastJobSummary := jobSummaries[0]
	if lastJobSummary.Status != expectedStatus {
		return nil, fmt.Errorf("method GetLastPipelineJobWithStatus for application %s expected status '%s', but it received '%s'",
			appName, expectedStatus, lastJobSummary.Status)
	}
	logger.Debug().Str("appName", appName).Str("expectedStatus", expectedStatus).Msg("method GetLastPipelineJobWithStatus for application received expected status")
	return lastJobSummary, nil
}

// GetAnyPipelineJobWithStatus Checks if any pipeline job exists with status
func GetAnyPipelineJobWithStatus(cfg config.Config, appName, expectedStatus string, logger zerolog.Logger) (*models.JobSummary, error) {
	jobSummaries, err := getPipelineJobs(cfg, appName)
	if err != nil {
		return nil, err
	}
	for _, jobSummary := range jobSummaries {
		if jobSummary.Status == expectedStatus {
			logger.Debug().Str("appName", appName).Str("expectedStatus", expectedStatus).Msg("method GetAnyPipelineJobWithStatus for application received expected status")
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
	client := httpUtils.GetJobClient(cfg)
	applicationJobs, err := client.GetApplicationJobs(params, nil)
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
	client := httpUtils.GetJobClient(cfg)
	params := pipelineJobClient.NewStopApplicationJobParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	jobStopped, err := client.StopApplicationJob(params, nil)
	if err == nil && jobStopped != nil {
		return nil
	}

	return fmt.Errorf("failed to stop job %s for an app %s. Error: %w", jobName, appName, err)
}

// IsDone Checks if job is done
func IsDone(cfg config.Config, appName, jobName string, logger zerolog.Logger) (string, error) {
	jobStatus, err := GetStatus(cfg, appName, jobName, logger)
	if err != nil {
		return "", err
	}
	if jobStatus == "Succeeded" || jobStatus == "Failed" {
		logger.Debug().Str("appName", appName).Str("jobName", jobName).Str("jobStatus", jobStatus).Msg("Job is done")
		return jobStatus, nil
	}
	logger.Debug().Str("appName", appName).Str("jobName", jobName).Msg("Job is not done yet")
	return "", fmt.Errorf("job %s for an app %s is not complete yet, Status %s", jobName, appName, jobStatus)
}

// GetStatus Gets status of job
func GetStatus(cfg config.Config, appName, jobName string, logger zerolog.Logger) (string, error) {
	job, err := Get(cfg, appName, jobName)
	if err != nil {
		return "", err
	}
	if job != nil {
		return job.Status, nil
	}
	logger.Debug().Str("appName", appName).Str("jobName", jobName).Msg("Job was not listed yet")
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

	client := httpUtils.GetJobClient(cfg)
	applicationJob, err := client.GetApplicationJob(params, nil)
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

	client := httpUtils.GetJobClient(cfg)
	applicationJob, err := client.GetApplicationJob(params, nil)
	if err == nil &&
		applicationJob.Payload != nil &&
		applicationJob.Payload.Steps != nil {

		return applicationJob.Payload.Steps
	}

	return nil
}

// GetLogForStep gets log for step
func GetLogForStep(cfg config.Config, appName, jobName, stepName string, logger zerolog.Logger) string {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := pipelineJobClient.NewGetPipelineJobStepLogsParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithStepName(stepName).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	client := httpUtils.GetJobClient(cfg)
	applicationJobLogs, err := client.GetPipelineJobStepLogs(params, nil)
	if err != nil {
		logger.Error().Str("appName", appName).Str("jobName", jobName).Str("stepName", stepName).Msg("failed to get pipeline log for the app")
		return ""
	}
	return applicationJobLogs.Payload
}
