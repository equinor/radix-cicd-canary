package build

import (
	"context"
	"strings"
	"time"

	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	jobUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const (
	Secret1            = "SECRET_1"
	Secret2            = "SECRET_2"
	Secret1Value       = "SECRET_1_VALUE"
	Secret2Value       = "SECRET_2_VALUE"
	Secret1ValueSha256 = "7cb08032ffb66e835ceeb10b849a8728440b0631ccb11f652b807534df26275e"
	Secret2ValueSha256 = "087f38fb04a52265ad5394fc20a6bfaa78c44bd58097dbcb690031a85b6e8313"
)

// Application Tests that we are able to successfully build an application
func Application(ctx context.Context, cfg config.Config) error {
	// Trigger build via web hook
	err := httpUtils.TriggerWebhookPush(ctx, cfg, defaults.App2BranchToBuildFrom, defaults.App2CommitID, defaults.App2SSHRepository, defaults.App2SharedSecret)
	if err != nil {
		return err
	}
	log.Ctx(ctx).Info().Msg("First job was triggered")

	// Get job
	jobSummary, err := test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
		return jobUtils.GetLastPipelineJobWithStatus(ctx, cfg, defaults.App2Name, "Running")
	})

	if err != nil {
		return err
	}

	jobName := *jobSummary.Name
	log.Ctx(ctx).Info().Msgf("First job name: %s", jobName)

	// Another build should cause second job to queue up
	// Trigger another build via web hook
	time.Sleep(1 * time.Second)
	err = httpUtils.TriggerWebhookPush(ctx, cfg, defaults.App2BranchToBuildFrom, defaults.App2CommitID, defaults.App2SSHRepository, defaults.App2SharedSecret)
	if err != nil {
		return err
	}
	log.Ctx(ctx).Info().Msg("Second job was triggered")

	err = test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		_, err := jobUtils.GetLastPipelineJobWithStatus(ctx, cfg, defaults.App2Name, "Queued")
		return err
	})

	if err != nil {
		return err
	}

	log.Ctx(ctx).Info().Msg("Second job was queued")
	jobStatus, err := test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (string, error) {
		return jobUtils.IsDone(ctx, cfg, defaults.App2Name, jobName)
	})
	if err != nil {
		return err
	}
	if jobStatus != "Succeeded" {
		return errors.Errorf("expected job status was Success, but got %s", jobStatus)
	}
	log.Ctx(ctx).Info().Msg("First job was completed")

	steps := jobUtils.GetSteps(ctx, cfg, defaults.App2Name, jobName)
	expectedSteps := jobUtils.NewExpectedSteps().
		Add("clone-config").
		Add("radix-pipeline").
		Add("clone", "app-qa").
		Add("clone", "app-prod").
		Add("clone", "redis-prod").
		Add("clone", "redis-qa").
		Add("build-app-qa", "app").
		Add("build-app-prod", "app").
		Add("build-redis-prod", "redis").
		Add("build-redis-qa", "redis")

	if len(steps) != expectedSteps.Count() {
		return errors.New("number of pipeline steps was not as expected")
	}

	for _, step := range steps {
		if !expectedSteps.HasStepWithComponent(step.Name, step.Components) {
			return errors.Errorf("missing expected step %s with components %s", step.Name, step.Components)
		}
	}

	log.Ctx(ctx).Debug().Str("jobName", jobName).Msg("Checking Sub-pipeline run...")
	pipelineRuns, err := jobUtils.GetPipelineRuns(ctx, cfg, defaults.App2Name, jobName)
	if err != nil {
		return err
	}
	run, ok := slice.FindFirst(pipelineRuns, func(run *models.PipelineRun) bool {
		return true
	})
	if !ok {
		return errors.New("No Pipeline run found")
	}

	tasks, err := jobUtils.GetPipelineRunTasks(ctx, cfg, defaults.App2Name, jobName, *run.KubeName)
	if err != nil {
		return err
	}
	targetTask, ok := slice.FindFirst(tasks, func(task *models.PipelineRunTask) bool {
		return *task.Name == "details"
	})
	if !ok {
		return errors.New("Tekton test task not found!")
	}

	// Test tekton log output contain parameters and secrets
	tektonLogContent, err := jobUtils.GetLogForPipelineStep(ctx, cfg, defaults.App2Name, jobName, *run.KubeName, *targetTask.KubeName, "test-tekton")
	if err != nil {
		return err
	}

	if !strings.Contains(tektonLogContent, Secret1Value) {
		return errors.New("Tekton test does not contain SecretValue")
	}

	if !strings.Contains(tektonLogContent, "github.com") {
		return errors.New("Tekton test does no conaint github.com (should be printed from known_hosts)")
	}
	log.Ctx(ctx).Info().Msg("Sub-pipeline completed")

	stepLog := jobUtils.GetLogForStep(ctx, cfg, defaults.App2Name, jobName, "build-app-qa")

	// Validate if Dockerfile build output contains SHA256 hash of build secrets:
	// https://github.com/equinor/radix-canarycicd-test-2/blob/master/Dockerfile#L9
	if !strings.Contains(stepLog, Secret1ValueSha256) || !strings.Contains(stepLog, Secret2ValueSha256) {
		return errors.New("build secrets are not contained in build log")
	}

	jobSummary, err = test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
		return jobUtils.GetLastPipelineJobWithStatus(ctx, cfg, defaults.App2Name, "Running")
	})

	if err != nil {
		return err
	}

	// Stop job and verify that it has been stopped
	jobName = *jobSummary.Name
	log.Ctx(ctx).Info().Msgf("Second job name: %s", jobName)
	err = jobUtils.Stop(ctx, cfg, defaults.App2Name, jobName)
	if err != nil {
		return err
	}

	err = test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		_, err := jobUtils.GetLastPipelineJobWithStatus(ctx, cfg, defaults.App2Name, "Stopped")
		return err
	})
	if err != nil {
		return err
	}

	log.Ctx(ctx).Info().Msg("Second job was stopped")

	return nil
}
