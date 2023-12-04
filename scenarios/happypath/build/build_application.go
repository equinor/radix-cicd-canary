package build

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/array"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
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

type expectedStep struct {
	name       string
	components []string
}

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
		return job.GetLastPipelineJobWithStatus(ctx, cfg, defaults.App2Name, "Running")
	})

	if err != nil {
		return err
	}

	jobName := jobSummary.Name
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
		_, err := job.GetLastPipelineJobWithStatus(ctx, cfg, defaults.App2Name, "Queued")
		return err
	})

	if err != nil {
		return err
	}

	log.Ctx(ctx).Info().Msg("Second job was queued")
	jobStatus, err := test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (string, error) {
		return job.IsDone(ctx, cfg, defaults.App2Name, jobName)
	})
	if err != nil {
		return err
	}
	if jobStatus != "Succeeded" {
		return errors.Errorf("expected job status was Success, but got %s", jobStatus)
	}
	log.Ctx(ctx).Info().Msg("First job was completed")
	steps := job.GetSteps(ctx, cfg, defaults.App2Name, jobName)

	expectedSteps := []expectedStep{
		{name: "clone-config", components: []string{}},
		{name: "prepare-pipelines", components: []string{}},
		{name: "radix-pipeline", components: []string{}},
		{name: "clone", components: []string{}},
		{name: "build-app", components: []string{"app"}},
		{name: "build-redis", components: []string{"redis"}},
		{name: "run-pipelines", components: []string{}},
	}

	if len(steps) != len(expectedSteps) {
		return errors.New("number of pipeline steps was not as expected")
	}

	for index, step := range steps {
		if !strings.EqualFold(step.Name, expectedSteps[index].name) {
			return errors.Errorf("expeced step %s, but got %s", expectedSteps[index].name, step.Name)
		}

		if !array.EqualElements(step.Components, expectedSteps[index].components) {
			return errors.Errorf("expeced components %s, but got %s", expectedSteps[index].components, step.Components)
		}
	}

	stepLog := job.GetLogForStep(ctx, cfg, defaults.App2Name, jobName, "build-app")
	// Validate if Dockerfile build output contains SHA256 hash of build secrets:
	// https://github.com/equinor/radix-canarycicd-test-2/blob/master/Dockerfile#L9
	if !strings.Contains(stepLog, Secret1ValueSha256) || !strings.Contains(stepLog, Secret2ValueSha256) {
		return errors.New("build secrets are not contained in build log")
	}

	jobSummary, err = test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
		return job.GetLastPipelineJobWithStatus(ctx, cfg, defaults.App2Name, "Running")
	})

	if err != nil {
		return err
	}

	// Stop job and verify that it has been stopped
	jobName = jobSummary.Name
	log.Ctx(ctx).Info().Msgf("Second job name: %s", jobName)
	err = job.Stop(ctx, cfg, defaults.App2Name, jobName)
	if err != nil {
		return err
	}

	err = test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		_, err := job.GetLastPipelineJobWithStatus(ctx, cfg, defaults.App2Name, "Stopped")
		return err
	})
	if err != nil {
		return err
	}

	log.Ctx(ctx).Info().Msg("Second job was stopped")

	// Test tekton log output contain parameters and secrets
	tektonLogContent := job.GetLogForStep(ctx, cfg, defaults.App2Name, jobName, "test-tekton")
	if !strings.Contains(tektonLogContent, Secret1Value) {
		return errors.New("Tekton test does not contain SecretValue")
	}

	if !strings.Contains(tektonLogContent, "github.com") {
		return errors.New("Tekton test does no conaint github.com (should be printed from known_hosts)")
	}

	return nil
}
