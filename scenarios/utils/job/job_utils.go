package job

import (
	"context"
	"fmt"

	pipelineJobClient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/pipeline_job"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// GetLastPipelineJobWithStatus Checks if a last pipeline job exists with status
func GetLastPipelineJobWithStatus(ctx context.Context, cfg config.Config, appName, expectedStatus string) (*models.JobSummary, error) {
	jobSummaries, err := getPipelineJobs(ctx, cfg, appName)
	if err != nil {
		return nil, err
	}
	lastJobSummary := jobSummaries[0]
	if lastJobSummary.Status != expectedStatus {
		return nil, errors.Errorf("method GetLastPipelineJobWithStatus for application %s expected status '%s', but it received '%s'",
			appName, expectedStatus, lastJobSummary.Status)
	}
	log.Ctx(ctx).Debug().Str("app", appName).Str("expectedStatus", expectedStatus).Msg("method GetLastPipelineJobWithStatus for application received expected status")
	return lastJobSummary, nil
}

// GetAnyPipelineJobWithStatus Checks if any pipeline job exists with status
func GetAnyPipelineJobWithStatus(ctx context.Context, cfg config.Config, appName, expectedStatus string) (*models.JobSummary, error) {
	jobSummaries, err := getPipelineJobs(ctx, cfg, appName)
	if err != nil {
		return nil, err
	}
	for _, jobSummary := range jobSummaries {
		if jobSummary.Status == expectedStatus {
			log.Ctx(ctx).Debug().Str("app", appName).Str("expectedStatus", expectedStatus).Msg("method GetAnyPipelineJobWithStatus for application received expected status")
			return jobSummary, nil
		}
	}
	return nil, errors.Errorf("method GetAnyPipelineJobWithStatus for application %s expected any job with the status '%s', but it does not exist",
		appName, expectedStatus)
}

func getPipelineJobs(ctx context.Context, cfg config.Config, appName string) ([]*models.JobSummary, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := pipelineJobClient.NewGetApplicationJobsParams().
		WithAppName(appName).
		WithContext(ctx).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)
	client := httpUtils.GetJobClient(cfg)
	applicationJobs, err := client.GetApplicationJobs(params, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error calling GetApplicationJobs for application %s", appName)
	}
	if len(applicationJobs.Payload) == 0 {
		return nil, errors.Errorf("method GetApplicationJobs for application %s received invalid or empty applicationJobs payload", appName)
	}
	jobSummaries := applicationJobs.Payload
	return jobSummaries, nil
}

// Stop Stops a job
func Stop(ctx context.Context, cfg config.Config, appName, jobName string) error {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()
	client := httpUtils.GetJobClient(cfg)
	params := pipelineJobClient.NewStopApplicationJobParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithContext(ctx).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	jobStopped, err := client.StopApplicationJob(params, nil)
	if err == nil && jobStopped != nil {
		return nil
	}

	return errors.Wrapf(err, "failed to stop job %s for an app %s", jobName, appName)
}

// IsDone Checks if job is done
func IsDone(ctx context.Context, cfg config.Config, appName, jobName string) (string, error) {
	jobStatus, err := GetStatus(ctx, cfg, appName, jobName)
	if err != nil {
		return "", err
	}
	if jobStatus == "Succeeded" || jobStatus == "Failed" {
		log.Ctx(ctx).Debug().Str("app", appName).Str("jobName", jobName).Str("jobStatus", jobStatus).Msg("Job is done")
		return jobStatus, nil
	}
	log.Ctx(ctx).Debug().Str("app", appName).Str("jobName", jobName).Msg("Job is not done yet")
	return "", errors.Errorf("job %s for an app %s is not complete yet, Status %s", jobName, appName, jobStatus)
}

// GetStatus Gets status of job
func GetStatus(ctx context.Context, cfg config.Config, appName, jobName string) (string, error) {
	job, err := Get(ctx, cfg, appName, jobName)
	if err != nil {
		return "", err
	}
	if job != nil {
		return job.Status, nil
	}
	log.Ctx(ctx).Debug().Str("app", appName).Str("jobName", jobName).Msg("Job was not listed yet")
	return "", errors.Errorf("job %s does not exist", jobName)
}

// Get gets job from job name
func Get(ctx context.Context, cfg config.Config, appName, jobName string) (*models.Job, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := pipelineJobClient.NewGetApplicationJobParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithContext(ctx).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	client := httpUtils.GetJobClient(cfg)
	applicationJob, err := client.GetApplicationJob(params, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if applicationJob.Payload != nil {
		return applicationJob.Payload, nil
	}
	return nil, errors.Errorf("failed to get job %s", jobName)
}

// GetSteps gets job from job name
func GetSteps(ctx context.Context, cfg config.Config, appName, jobName string) []*models.Step {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := pipelineJobClient.NewGetApplicationJobParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithContext(ctx).
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
func GetLogForStep(ctx context.Context, cfg config.Config, appName, jobName, stepName string) string {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := pipelineJobClient.NewGetPipelineJobStepLogsParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithStepName(stepName).
		WithContext(ctx).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	client := httpUtils.GetJobClient(cfg)
	applicationJobLogs, err := client.GetPipelineJobStepLogs(params, nil)
	if err != nil {
		log.Ctx(ctx).Error().Str("jobName", jobName).Str("stepName", stepName).Msg("failed to get pipeline log for the app")
		return ""
	}
	return applicationJobLogs.Payload
}

// GetPipelineRuns gets pipeline runs from job name
func GetPipelineRuns(ctx context.Context, cfg config.Config, appName, jobName string) ([]*models.PipelineRun, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := pipelineJobClient.NewGetTektonPipelineRunsParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithContext(ctx).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	client := httpUtils.GetJobClient(cfg)
	response, err := client.GetTektonPipelineRuns(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pipelinerun details: %w", err)
	}

	return response.Payload, nil
}

// GetPipelineRunTasks gets tasks from pipeline run from job name
func GetPipelineRunTasks(ctx context.Context, cfg config.Config, appName, jobName, pipelineRunName string) ([]*models.PipelineRunTask, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := pipelineJobClient.NewGetTektonPipelineRunTasksParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithPipelineRunName(pipelineRunName).
		WithContext(ctx).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	client := httpUtils.GetJobClient(cfg)
	applicationJob, err := client.GetTektonPipelineRunTasks(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tasks for pipelinerun: %w", err)
	}

	return applicationJob.Payload, nil
}

func GetLogForPipelineStep(ctx context.Context, cfg config.Config, appName, jobName, pipelineRunName, taskName, stepName string) (string, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := pipelineJobClient.NewGetTektonPipelineRunTaskStepLogsParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithPipelineRunName(pipelineRunName).
		WithTaskName(taskName).
		WithStepName(stepName).
		WithContext(ctx).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	client := httpUtils.GetJobClient(cfg)
	logs, err := client.GetTektonPipelineRunTaskStepLogs(params, nil)
	if err != nil {
		return "", fmt.Errorf("failed to fetch logs for pipeline task-step log: %w", err)
	}
	return logs.Payload, nil
}
