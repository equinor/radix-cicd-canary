package build

import (
	"context"
	"fmt"
	"strings"
	"time"

	"errors"

	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httputils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	jobutils "github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/equinor/radix-common/utils/slice"
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
	sharedSecret, err := application.GetSharedSecret(ctx, cfg, defaults.App2Name)
	if err != nil {
		return err
	}

	// Trigger build via web hook
	err = httputils.TriggerWebhookPush(ctx, cfg, defaults.App2BranchToBuildFrom, defaults.App2CommitID, defaults.App2SSHRepository, sharedSecret)
	if err != nil {
		return err
	}
	log.Ctx(ctx).Info().Msg("First job was triggered")

	// Get job
	jobSummary, err := test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
		return jobutils.GetLastPipelineJobWithStatus(ctx, cfg, defaults.App2Name, "Running")
	})

	if err != nil {
		return err
	}

	jobName := *jobSummary.Name
	log.Ctx(ctx).Info().Msgf("First job name: %s", jobName)

	// Another build should cause second job to queue up
	// Trigger another build via web hook
	time.Sleep(1 * time.Second)
	err = httputils.TriggerWebhookPush(ctx, cfg, defaults.App2BranchToBuildFrom, defaults.App2CommitID, defaults.App2SSHRepository, sharedSecret)
	if err != nil {
		return err
	}
	log.Ctx(ctx).Info().Msg("Second job was triggered")

	err = test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		_, err := jobutils.GetLastPipelineJobWithStatus(ctx, cfg, defaults.App2Name, "Queued")
		return err
	})

	if err != nil {
		return err
	}

	log.Ctx(ctx).Info().Msg("Second job was queued")
	jobStatus, err := test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (string, error) {
		return jobutils.IsDone(ctx, cfg, defaults.App2Name, jobName)
	})
	if err != nil {
		return err
	}
	if jobStatus != "Succeeded" {
		return fmt.Errorf("expected job status was Success, but got %s", jobStatus)
	}
	log.Ctx(ctx).Info().Msg("First job was completed")

	steps := jobutils.GetSteps(ctx, cfg, defaults.App2Name, jobName)
	expectedSteps := jobutils.NewExpectedSteps().
		Add(jobutils.StepTypeCloneConfig).
		Add(jobutils.StepTypePipelineRunner).
		Add(jobutils.StepTypeClone, jobutils.WithComponents("app"), jobutils.WithEnvironments("qa")).
		Add(jobutils.StepTypeClone, jobutils.WithComponents("app"), jobutils.WithEnvironments("prod")).
		Add(jobutils.StepTypeClone, jobutils.WithComponents("redis"), jobutils.WithEnvironments("prod")).
		Add(jobutils.StepTypeClone, jobutils.WithComponents("redis"), jobutils.WithEnvironments("qa")).
		Add(jobutils.StepTypeSubPipelineStep,
			jobutils.WithStatus(jobutils.StepStatusCompleted),
			jobutils.WithSubPipelineTaskStep(&models.SubPipelineTaskStep{
				Environment:  new("qa"),
				PipelineName: new("radix-cicdcanary-test2")})).
		Add(jobutils.StepTypeSubPipelineStep,
			jobutils.WithStatus(jobutils.StepStatusCompleted),
			jobutils.WithSubPipelineTaskStep(&models.SubPipelineTaskStep{
				Environment:  new("prod"),
				PipelineName: new("radix-cicdcanary-test2")})).
		Add(jobutils.StepTypeBuild, jobutils.WithComponents("app"), jobutils.WithEnvironments("qa")).
		Add(jobutils.StepTypeBuild, jobutils.WithComponents("app"), jobutils.WithEnvironments("prod")).
		Add(jobutils.StepTypeBuild, jobutils.WithComponents("redis"), jobutils.WithEnvironments("prod")).
		Add(jobutils.StepTypeBuild, jobutils.WithComponents("redis"), jobutils.WithEnvironments("qa"))

	if len(steps) != expectedSteps.Count() {
		return errors.New("number of pipeline steps was not as expected")
	}

	for _, step := range steps {
		if !expectedSteps.HasStep(*step) {
			return fmt.Errorf("step %s not found in list of expected steps", *step.Name)
		}
	}

	logStep, ok := slice.FindFirst(steps, func(step *models.Step) bool {
		if step == nil {
			return false
		}
		return step.Type == string(jobutils.StepTypeBuild) && slice.ElementsMatch(step.Components, []string{"app"}) && slice.ElementsMatch(step.Environments, []string{"qa"})
	})

	if !ok {
		return errors.New("failed to find build step for app in qa environment")
	}

	stepLog := jobutils.GetLogForStep(ctx, cfg, defaults.App2Name, jobName, *logStep.Name)

	// Validate if Dockerfile build output contains SHA256 hash of build secrets:
	// https://github.com/equinor/radix-canarycicd-test-2/blob/master/Dockerfile#L9
	if !strings.Contains(stepLog, Secret1ValueSha256) || !strings.Contains(stepLog, Secret2ValueSha256) {
		return errors.New("build secrets are not contained in build log")
	}

	log.Ctx(ctx).Debug().Str("jobName", jobName).Msg("Checking Sub-pipeline run...")
	pipelineRuns, err := jobutils.GetPipelineRuns(ctx, cfg, defaults.App2Name, jobName)
	if err != nil {
		return err
	}
	run, ok := slice.FindFirst(pipelineRuns, func(run *models.PipelineRun) bool {
		return true
	})
	if !ok {
		return errors.New("failed to find pipeline runs")
	}

	tasks, err := jobutils.GetPipelineRunTasks(ctx, cfg, defaults.App2Name, jobName, *run.KubeName)
	if err != nil {
		return err
	}
	targetTask, ok := slice.FindFirst(tasks, func(task *models.PipelineRunTask) bool {
		return *task.Name == "details"
	})
	if !ok {
		return errors.New("failed to find Tekton task")
	}

	// Test tekton log output contain parameters and secrets
	tektonLogContent, err := jobutils.GetLogForPipelineStep(ctx, cfg, defaults.App2Name, jobName, *run.KubeName, *targetTask.KubeName, "test-tekton")
	if err != nil {
		return err
	}

	if !strings.Contains(tektonLogContent, Secret1Value) {
		return errors.New("tekton test does not contain SecretValue")
	}

	if !strings.Contains(tektonLogContent, "github.com") {
		return errors.New("tekton test does not contain github.com (should be printed from known_hosts)")
	}
	log.Ctx(ctx).Info().Msg("Sub-pipeline completed")

	jobSummary, err = test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
		return jobutils.GetLastPipelineJobWithStatus(ctx, cfg, defaults.App2Name, "Running")
	})

	if err != nil {
		return err
	}

	// Stop job and verify that it has been stopped
	jobName = *jobSummary.Name
	log.Ctx(ctx).Info().Msgf("Second job name: %s", jobName)
	err = jobutils.Stop(ctx, cfg, defaults.App2Name, jobName)
	if err != nil {
		return err
	}

	err = test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		_, err := jobutils.GetLastPipelineJobWithStatus(ctx, cfg, defaults.App2Name, "Stopped")
		return err
	})
	if err != nil {
		return err
	}

	log.Ctx(ctx).Info().Msg("Second job was stopped")

	return nil
}
