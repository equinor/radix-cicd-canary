package buildsecrets

import (
	"context"
	"fmt"
	"strings"

	applicationClient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/build"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/rs/zerolog/log"
)

// Set Tests that we are able to successfully set build secrets
func Set(ctx context.Context, cfg config.Config, suiteName string) error {
	appName := defaults.App2Name
	appCtx := log.Ctx(ctx).With().Str("app", appName).Logger().WithContext(ctx)
	// Trigger build to apply RA with build secrets
	err := httpUtils.TriggerWebhookPush(cfg, defaults.App2BranchToBuildFrom, defaults.App2CommitID, defaults.App2SSHRepository, defaults.App2SharedSecret, appCtx)
	if err != nil {
		return err
	}

	log.Ctx(appCtx).Info().Msg("Job was triggered to apply RA")

	// Get job
	jobSummary, err := test.WaitForCheckFuncWithValueOrTimeout(appCtx, cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
		return job.GetLastPipelineJobWithStatus(ctx, cfg, appName, "Failed")
	})
	if err != nil {
		return err
	}

	jobName := jobSummary.Name
	job, err := job.Get(ctx, cfg, appName, jobName)
	if err != nil {
		return err
	}
	expectedSteps := []string{
		"clone-config",
		"prepare-pipelines",
		"radix-pipeline"}

	if len(job.Steps) != len(expectedSteps) {
		return fmt.Errorf("job should not contain any build step")
	}

	// First job failed, due to missing build secrets, as expected in test
	// Set build secrets
	err = test.WaitForCheckFuncOrTimeout(appCtx, cfg, func(cfg config.Config, ctx context.Context) error {
		return buildSecretsAreListedWithStatus(ctx, cfg, appName, "Pending")
	})

	if err != nil {
		return err
	}

	err = setSecret(appCtx, cfg, appName, build.Secret1, build.Secret1Value)
	if err != nil {
		return err
	}

	err = setSecret(appCtx, cfg, appName, build.Secret2, build.Secret2Value)
	if err != nil {
		return err
	}

	return test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		return buildSecretsAreListedWithStatus(ctx, cfg, appName, "Consistent")
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
	return fmt.Errorf("failed buildSecretsAreListedWithStatus expected %s", expectedStatus)
}

func setSecret(ctx context.Context, cfg config.Config, appName, secretName, secretValue string) error {
	log.Ctx(ctx).Debug().Msgf("setSecret %s with value %s", secretName, secretValue)
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
		return fmt.Errorf("failed to set secret %s. Error: %v", secretName, err)
	}

	return nil
}
