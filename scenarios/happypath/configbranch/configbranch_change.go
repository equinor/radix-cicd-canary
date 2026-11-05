package configbranch

import (
	"context"
	"fmt"

	"errors"

	apiclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httputils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	jobutils "github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/rs/zerolog/log"
)

// Change Tests that radixconfig is read from the branch defined as configBranch
func Change(ctx context.Context, cfg config.Config) error {
	appName := defaults.App4Name

	sharedSecret, err := application.GetSharedSecret(ctx, cfg, appName)
	if err != nil {
		return err
	}

	// Trigger first build via web hook
	err = httputils.TriggerWebhookPush(ctx, cfg, defaults.App4ConfigBranch, defaults.App4CommitID, defaults.App4SSHRepository, sharedSecret)
	if err != nil {
		return err
	}

	log.Ctx(ctx).Info().Msgf("First job was triggered")
	jobSummary, err := waitForJobRunning(ctx, cfg, appName)

	if err != nil {
		return fmt.Errorf("first job for application %s: %w", defaults.App4Name, err)
	}

	jobName := *jobSummary.Name
	log.Ctx(ctx).Info().Msgf("First job name: %s", jobName)

	if err = waitForJobDone(ctx, cfg, appName, jobName); err != nil {
		return fmt.Errorf("first job for application %s: %w", defaults.App4Name, err)
	}

	log.Ctx(ctx).Info().Msg("First job was completed")

	expectedSteps := jobutils.NewExpectedSteps().
		Add(jobutils.StepTypeCloneConfig).
		Add(jobutils.StepTypePipelineRunner).
		Add(jobutils.StepTypeClone, jobutils.WithComponents("www"), jobutils.WithEnvironments("prod")).
		Add(jobutils.StepTypeBuild, jobutils.WithComponents("www"), jobutils.WithEnvironments("prod"))

	if ok, err := validateJobSteps(ctx, cfg, appName, jobName, expectedSteps); !ok {
		return err
	}

	// Change config branch, trigger second webhook and verify job
	if err := patchConfigBranch(ctx, cfg, appName, defaults.App4NewConfigBranch); err != nil {
		return err
	}

	err = httputils.TriggerWebhookPush(ctx, cfg, defaults.App4NewConfigBranch, defaults.App4NewCommitID, defaults.App4SSHRepository, sharedSecret)
	if err != nil {
		return err
	}

	log.Ctx(ctx).Info().Msg("Second job was triggered")
	jobSummary, err = waitForJobRunning(ctx, cfg, appName)

	if err != nil {
		return fmt.Errorf("second job for application %s: %w", defaults.App4Name, err)
	}

	jobName = *jobSummary.Name
	log.Ctx(ctx).Info().Str("jobName", jobName).Msg("Second job name")

	if err = waitForJobDone(ctx, cfg, appName, jobName); err != nil {
		return fmt.Errorf("second job for application %s: %w", defaults.App4Name, err)
	}

	log.Ctx(ctx).Info().Msg("Second job was completed")

	expectedSteps = jobutils.NewExpectedSteps().
		Add(jobutils.StepTypeCloneConfig).
		Add(jobutils.StepTypePipelineRunner).
		Add(jobutils.StepTypeClone, jobutils.WithComponents("www2"), jobutils.WithEnvironments("prod")).
		Add(jobutils.StepTypeBuild, jobutils.WithComponents("www2"), jobutils.WithEnvironments("prod"))

	if ok, err := validateJobSteps(ctx, cfg, appName, jobName, expectedSteps); !ok {
		return err
	}

	return nil
}

func waitForJobRunning(ctx context.Context, cfg config.Config, appName string) (*models.JobSummary, error) {
	status := "Running"

	return test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
		return jobutils.GetLastPipelineJobWithStatus(ctx, cfg, appName, status)
	})
}

func waitForJobDone(ctx context.Context, cfg config.Config, appName, jobName string) error {
	jobStatus, err := test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (string, error) {
		return jobutils.IsDone(ctx, cfg, appName, jobName)
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
	log.Ctx(ctx).Debug().Str("app", appName).Msgf("Set ConfigBranch to %v", newConfigBranch)
	patchRequest := models.ApplicationRegistrationPatchRequest{
		ApplicationRegistrationPatch: &models.ApplicationRegistrationPatch{
			ConfigBranch: newConfigBranch,
		},
	}

	params := apiclient.NewModifyRegistrationDetailsParams().
		WithAppName(appName).
		WithPatchRequest(&patchRequest)

	client := httputils.GetApplicationClient(cfg)
	_, err := client.ModifyRegistrationDetailsContext(ctx, params, nil)
	if err != nil {
		return err
	}
	log.Ctx(ctx).Debug().Str("app", appName).Msgf("ConfigBranch has been set to %v", newConfigBranch)
	return nil
}

func validateJobSteps(ctx context.Context, cfg config.Config, appName, jobName string, expectedSteps *jobutils.ExpectedSteps) (bool, error) {
	steps := jobutils.GetSteps(ctx, cfg, appName, jobName)

	if len(steps) != expectedSteps.Count() {
		return false, errors.New("number of pipeline steps was not as expected")
	}

	for _, step := range steps {
		if !expectedSteps.HasStep(*step) {
			return false, fmt.Errorf("step %s not found in list of expected steps", *step.Name)
		}
	}
	return true, nil
}
