package configbranch

import (
	"context"
	"fmt"
	"strings"

	apiclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/array"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type expectedStep struct {
	name       string
	components []string
}

// Change Tests that radixconfig is read from the branch defined as configBranch
func Change(ctx context.Context, cfg config.Config) error {
	appName := defaults.App4Name

	// Trigger first build via web hook
	err := httpUtils.TriggerWebhookPush(ctx, cfg, defaults.App4ConfigBranch, defaults.App4CommitID, defaults.App4SSHRepository, defaults.App4SharedSecret)
	if err != nil {
		return err
	}

	log.Ctx(ctx).Info().Msgf("First job was triggered")
	jobSummary, err := waitForJobRunning(ctx, cfg, appName)

	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("first job for application %s", defaults.App4Name))
	}

	jobName := jobSummary.Name
	log.Ctx(ctx).Info().Str("jobName", jobName).Msg("First job name")

	if err = waitForJobDone(ctx, cfg, appName, jobName); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("first job for application %s", defaults.App4Name))
	}

	log.Ctx(ctx).Info().Msg("First job was completed")

	expectedSteps := []expectedStep{
		{name: "clone-config", components: []string{}},
		{name: "prepare-pipelines", components: []string{}},
		{name: "radix-pipeline", components: []string{}},
		{name: "clone", components: []string{}},
		{name: "build-www", components: []string{"www"}},
		// {name: "run-pipelines", components: []string{}},//skip due to there is no sub-pipeline
	}

	if ok, err := validateJobSteps(ctx, cfg, jobName, appName, expectedSteps); !ok {
		return err
	}

	// Change config branch, trigger second webhook and verify job
	if err := patchConfigBranch(ctx, cfg, appName, defaults.App4NewConfigBranch); err != nil {
		return err
	}

	err = httpUtils.TriggerWebhookPush(ctx, cfg, defaults.App4NewConfigBranch, defaults.App4NewCommitID, defaults.App4SSHRepository, defaults.App4SharedSecret)
	if err != nil {
		return err
	}

	log.Ctx(ctx).Info().Msg("Second job was triggered")
	jobSummary, err = waitForJobRunning(ctx, cfg, appName)

	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("second job for application %s", defaults.App4Name))
	}

	jobName = jobSummary.Name
	log.Ctx(ctx).Info().Str("jobName", jobName).Msg("Second job name")

	if err = waitForJobDone(ctx, cfg, appName, jobName); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("second job for application %s", defaults.App4Name))
	}

	log.Ctx(ctx).Info().Msg("Second job was completed")

	expectedSteps = []expectedStep{
		{name: "clone-config", components: []string{}},
		{name: "prepare-pipelines", components: []string{}},
		{name: "radix-pipeline", components: []string{}},
		{name: "clone", components: []string{}},
		{name: "build-www2", components: []string{"www2"}},
		// {name: "run-pipelines", components: []string{}},//skip due to there is no sub-pipeline
	}

	if ok, err := validateJobSteps(ctx, cfg, appName, jobName, expectedSteps); !ok {
		return err
	}

	return nil
}

func waitForJobRunning(ctx context.Context, cfg config.Config, appName string) (*models.JobSummary, error) {
	status := "Running"

	return test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
		return job.GetLastPipelineJobWithStatus(ctx, cfg, appName, status)
	})
}

func waitForJobDone(ctx context.Context, cfg config.Config, appName, jobName string) error {
	jobStatus, err := test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (string, error) {
		return job.IsDone(cfg, appName, jobName, ctx)
	})
	if err != nil {
		return err
	}
	if jobStatus != "Succeeded" {
		return fmt.Errorf("job %s completed with status %s", jobName, jobStatus)
	}
	return nil
}

func patchConfigBranch(ctx context.Context, cfg config.Config, appName, newConfigBranch string) error {
	log.Ctx(ctx).Debug().Msgf("Set ConfigBranch to %v", newConfigBranch)
	patchRequest := models.ApplicationRegistrationPatchRequest{
		ApplicationRegistrationPatch: &models.ApplicationRegistrationPatch{
			ConfigBranch: newConfigBranch,
		},
	}

	params := apiclient.NewModifyRegistrationDetailsParams().
		WithAppName(appName).
		WithContext(ctx).
		WithPatchRequest(&patchRequest)

	client := httpUtils.GetApplicationClient(cfg)
	_, err := client.ModifyRegistrationDetails(params, nil)
	if err != nil {
		return err
	}
	log.Ctx(ctx).Debug().Msgf("ConfigBranch has been set to %v", newConfigBranch)
	return nil
}

func validateJobSteps(ctx context.Context, cfg config.Config, appName, jobName string, expectedSteps []expectedStep) (bool, error) {
	steps := job.GetSteps(ctx, cfg, appName, jobName)

	if len(steps) != len(expectedSteps) {
		return false, fmt.Errorf("number of pipeline steps was not as expected")
	}

	for index, step := range steps {
		if !strings.EqualFold(step.Name, expectedSteps[index].name) {
			return false, fmt.Errorf("expeced step %s, but got %s", expectedSteps[index].name, step.Name)
		}

		if !array.EqualElements(step.Components, expectedSteps[index].components) {
			return false, fmt.Errorf("expeced components %s, but got %s", expectedSteps[index].components, step.Components)
		}
	}

	return true, nil
}
