package buildsecrets

import (
	"context"
	"strings"

	applicationClient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/build"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	jobUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Set Tests that we are able to successfully set build secrets
func Set(ctx context.Context, cfg config.Config) error {
	// Trigger build to apply RA with build secrets
	err := httpUtils.TriggerWebhookPush(ctx, cfg, defaults.App2BranchToBuildFrom, defaults.App2CommitID, defaults.App2SSHRepository, defaults.App2SharedSecret)
	if err != nil {
		return err
	}

	log.Ctx(ctx).Info().Msg("Job was triggered to apply RA")

	// Get job
	jobSummary, err := test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
		return jobUtils.GetLastPipelineJobWithStatus(ctx, cfg, defaults.App2Name, "Failed")
	})
	if err != nil {
		return err
	}

	jobName := *jobSummary.Name
	job, err := jobUtils.Get(ctx, cfg, defaults.App2Name, jobName)
	if err != nil {
		return err
	}

	expectedSteps := jobUtils.NewExpectedSteps().
		Add("clone-config").
		Add("prepare-pipelines").
		Add("radix-pipeline")

	if len(job.Steps) != expectedSteps.Count() {
		return errors.New("job should not contain any build step")
	}

	// First job failed, due to missing build secrets, as expected in test
	// Set build secrets
	err = test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		return buildSecretsAreListedWithStatus(ctx, cfg, defaults.App2Name, "Pending")
	})

	if err != nil {
		return err
	}

	err = setSecret(ctx, cfg, defaults.App2Name, build.Secret1, build.Secret1Value)
	if err != nil {
		return err
	}

	err = setSecret(ctx, cfg, defaults.App2Name, build.Secret2, build.Secret2Value)
	if err != nil {
		return err
	}

	return test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		return buildSecretsAreListedWithStatus(ctx, cfg, defaults.App2Name, "Consistent")
	})
}

func buildSecretsAreListedWithStatus(ctx context.Context, cfg config.Config, appName, expectedStatus string) error {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := applicationClient.NewGetBuildSecretsParams().
		WithAppName(appName).
		WithContext(ctx).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)
	client := httpUtils.GetApplicationClient(cfg)
	buildSecrets, err := client.GetBuildSecrets(params, nil)
	if err == nil && buildSecrets.Payload != nil && len(buildSecrets.Payload) == 2 {
		if strings.EqualFold(*buildSecrets.Payload[0].Name, build.Secret1) &&
			strings.EqualFold(buildSecrets.Payload[0].Status, expectedStatus) &&
			strings.EqualFold(*buildSecrets.Payload[1].Name, build.Secret2) &&
			strings.EqualFold(buildSecrets.Payload[1].Status, expectedStatus) {
			return nil
		}
	}

	log.Ctx(ctx).Info().Msg("Build secrets are not listed yet")
	return errors.Errorf("failed buildSecretsAreListedWithStatus expected %s", expectedStatus)
}

func setSecret(ctx context.Context, cfg config.Config, appName, secretName, secretValue string) error {
	log.Ctx(ctx).Debug().Str("app", appName).Msgf("setSecret %s with value %s", secretName, secretValue)
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	secretParameters := models.SecretParameters{
		SecretValue: &secretValue,
	}

	params := applicationClient.NewUpdateBuildSecretsSecretValueParams().
		WithAppName(appName).
		WithContext(ctx).
		WithSecretName(secretName).
		WithSecretValue(&secretParameters).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	client := httpUtils.GetApplicationClient(cfg)
	_, err := client.UpdateBuildSecretsSecretValue(params, nil)
	if err != nil {
		return errors.Errorf("failed to set secret %s. Error: %v", secretName, err)
	}

	return nil
}
