package configbranch

import (
	"context"

	apiclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	jobUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

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
		return errors.Wrapf(err, "first job for application %s", defaults.App4Name)
	}

	jobName := jobSummary.Name
	log.Ctx(ctx).Info().Msgf("First job name: %s", jobName)

	if err = waitForJobDone(ctx, cfg, appName, jobName); err != nil {
		return errors.Wrapf(err, "first job for application %s", defaults.App4Name)
	}

	log.Ctx(ctx).Info().Msg("First job was completed")

	expectedSteps := jobUtils.NewExpectedSteps().
		Add("clone-config").
		Add("prepare-pipelines").
		Add("radix-pipeline").
		Add("clone", "www-prod").
		Add("build-www-prod", "www")

	if ok, err := validateJobSteps(ctx, cfg, appName, jobName, expectedSteps); !ok {
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
		return errors.Wrapf(err, "second job for application %s", defaults.App4Name)
	}

	jobName = jobSummary.Name
	log.Ctx(ctx).Info().Str("jobName", jobName).Msg("Second job name")

	if err = waitForJobDone(ctx, cfg, appName, jobName); err != nil {
		return errors.Wrapf(err, "second job for application %s", defaults.App4Name)
	}

	log.Ctx(ctx).Info().Msg("Second job was completed")

	expectedSteps = jobUtils.NewExpectedSteps().
		Add("clone-config").
		Add("prepare-pipelines").
		Add("radix-pipeline").
		Add("clone", "www2-prod").
		Add("build-www2-prod", "www2")

	if ok, err := validateJobSteps(ctx, cfg, appName, jobName, expectedSteps); !ok {
		return err
	}

	return nil
}

func waitForJobRunning(ctx context.Context, cfg config.Config, appName string) (*models.JobSummary, error) {
	status := "Running"

	return test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
		return jobUtils.GetLastPipelineJobWithStatus(ctx, cfg, appName, status)
	})
}

func waitForJobDone(ctx context.Context, cfg config.Config, appName, jobName string) error {
	jobStatus, err := test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (string, error) {
		return jobUtils.IsDone(ctx, cfg, appName, jobName)
	})
	if err != nil {
		return err
	}
	if jobStatus != "Succeeded" {
		return errors.Errorf("job %s completed with status %s", jobName, jobStatus)
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
		WithContext(ctx).
		WithPatchRequest(&patchRequest)

	client := httpUtils.GetApplicationClient(cfg)
	_, err := client.ModifyRegistrationDetails(params, nil)
	if err != nil {
		return err
	}
	log.Ctx(ctx).Debug().Str("app", appName).Msgf("ConfigBranch has been set to %v", newConfigBranch)
	return nil
}

func validateJobSteps(ctx context.Context, cfg config.Config, appName, jobName string, expectedSteps jobUtils.ExpectedSteps) (bool, error) {
	steps := jobUtils.GetSteps(ctx, cfg, appName, jobName)

	if len(steps) != expectedSteps.Count() {
		return false, errors.New("number of pipeline steps was not as expected")
	}

	for _, step := range steps {
		if !expectedSteps.HasStepWithComponent(step.Name, step.Components) {
			return false, errors.Errorf("missing expected step %s with components %s", step.Name, step.Components)
		}
	}
	return true, nil
}
