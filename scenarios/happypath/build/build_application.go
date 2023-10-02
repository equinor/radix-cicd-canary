package build

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

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
func Application(ctx context.Context, cfg config.Config, suiteName string) error {
	appName := defaults.App2Name
	appCtx := log.Ctx(ctx).With().Str("app", appName).Logger().WithContext(ctx)

	// Trigger build via web hook
	err := httpUtils.TriggerWebhookPush(cfg, defaults.App2BranchToBuildFrom, defaults.App2CommitID, defaults.App2SSHRepository, defaults.App2SharedSecret, appCtx)
	if err != nil {
		return err
	}
	log.Ctx(appCtx).Info().Msg("First job was triggered")

	// Get job
	jobSummary, err := test.WaitForCheckFuncWithValueOrTimeout(appCtx, cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
		return job.GetLastPipelineJobWithStatus(ctx, cfg, appName, "Running")
	})

	if err != nil {
		return err
	}

	jobName := jobSummary.Name
	firstJobCtx := log.Ctx(appCtx).With().Str("job", jobName).Logger().WithContext(appCtx)
	log.Ctx(firstJobCtx).Info().Msg("First job name")

	// Another build should cause second job to queue up
	// Trigger another build via web hook
	time.Sleep(1 * time.Second)
	err = httpUtils.TriggerWebhookPush(cfg, defaults.App2BranchToBuildFrom, defaults.App2CommitID, defaults.App2SSHRepository, defaults.App2SharedSecret, appCtx)
	if err != nil {
		return err
	}
	log.Ctx(appCtx).Info().Msg("Second job was triggered")

	err = test.WaitForCheckFuncOrTimeout(appCtx, cfg, func(cfg config.Config, ctx context.Context) error {
		_, err := job.GetLastPipelineJobWithStatus(ctx, cfg, appName, "Queued")
		return err
	})

	if err != nil {
		return err
	}

	log.Ctx(appCtx).Info().Msg("Second job was queued")
	jobStatus, err := test.WaitForCheckFuncWithValueOrTimeout(appCtx, cfg, func(cfg config.Config, ctx context.Context) (string, error) {
		return job.IsDone(cfg, appName, jobName, ctx)
	})
	if err != nil {
		return err
	}
	if jobStatus != "Succeeded" {
		return fmt.Errorf("expected job status was Success, but got %s", jobStatus)
	}
	log.Ctx(appCtx).Info().Msg("First job was completed")
	steps := job.GetSteps(appCtx, cfg, appName, jobName)

	expectedSteps := []expectedStep{
		{name: "clone-config", components: []string{}},
		{name: "prepare-pipelines", components: []string{}},
		{name: "radix-pipeline", components: []string{}},
		{name: "clone", components: []string{}},
		{name: "build-app", components: []string{"app"}},
		{name: "build-redis", components: []string{"redis"}},
	}

	if len(steps) != len(expectedSteps) {
		return errors.New("number of pipeline steps was not as expected")
	}

	for index, step := range steps {
		if !strings.EqualFold(step.Name, expectedSteps[index].name) {
			return fmt.Errorf("expeced step %s, but got %s", expectedSteps[index].name, step.Name)
		}

		if !array.EqualElements(step.Components, expectedSteps[index].components) {
			return fmt.Errorf("expeced components %s, but got %s", expectedSteps[index].components, step.Components)
		}
	}

	stepLog := job.GetLogForStep(appCtx, cfg, appName, jobName, "build-app")
	// Validate if Dockerfile build output contains SHA256 hash of build secrets:
	// https://github.com/equinor/radix-canarycicd-test-2/blob/master/Dockerfile#L9
	if !strings.Contains(stepLog, Secret1ValueSha256) || !strings.Contains(stepLog, Secret2ValueSha256) {
		return errors.New("build secrets are not contained in build log")
	}

	jobSummary, err = test.WaitForCheckFuncWithValueOrTimeout(appCtx, cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
		return job.GetLastPipelineJobWithStatus(ctx, cfg, appName, "Running")
	})

	if err != nil {
		return err
	}

	// Stop job and verify that it has been stopped
	jobName = jobSummary.Name
	jobCtx := log.Ctx(appCtx).With().Str("job", jobName).Logger().WithContext(appCtx)
	log.Ctx(jobCtx).Info().Msg("Second job name")
	err = job.Stop(jobCtx, cfg, appName, jobName)
	if err != nil {
		return err
	}

	err = test.WaitForCheckFuncOrTimeout(jobCtx, cfg, func(cfg config.Config, ctx context.Context) error {
		_, err := job.GetLastPipelineJobWithStatus(ctx, cfg, appName, "Stopped")
		return err
	})
	if err != nil {
		return err
	}

	log.Ctx(jobCtx).Info().Msg("Second job was stopped")
	return nil
}
