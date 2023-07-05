package buildsecrets

import (
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
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Set Tests that we are able to successfully set build secrets
func Set(cfg config.Config, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Trigger build to apply RA with build secrets
	err := httpUtils.TriggerWebhookPush(cfg, defaults.App2BranchToBuildFrom, defaults.App2CommitID, defaults.App2SSHRepository, defaults.App2SharedSecret, logger)
	if err != nil {
		return err
	}

	logger.Info("Job was triggered to apply RA")

	// Get job
	jobSummary, err := test.WaitForCheckFuncWithValueOrTimeout(cfg, func(cfg config.Config) (*models.JobSummary, error) {
		return job.IsListedWithStatus(cfg, defaults.App2Name, "Failed", logger)
	}, logger)
	if err != nil {
		return err
	}

	jobName := jobSummary.Name
	job, _ := job.Get(cfg, defaults.App2Name, jobName)

	expectedSteps := []string{
		"clone-config",
		"prepare-pipelines",
		"radix-pipeline"}

	if len(job.Steps) != len(expectedSteps) {
		return fmt.Errorf("job should not contain any build step")
	}

	// First job failed, due to missing build secrets, as expected in test
	// Set build secrets
	err = test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return buildSecretsAreListedWithStatus(cfg, "Pending")
	}, logger)

	if err != nil {
		return err
	}

	err = setSecret(cfg, build.Secret1, build.Secret1Value)
	if err != nil {
		return err
	}

	err = setSecret(cfg, build.Secret2, build.Secret2Value)
	if err != nil {
		return err
	}

	return test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return buildSecretsAreListedWithStatus(cfg, "Consistent")
	}, logger)
}

func buildSecretsAreListedWithStatus(cfg config.Config, expectedStatus string) error {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := applicationClient.NewGetBuildSecretsParams().
		WithAppName(defaults.App2Name).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetApplicationClient(cfg)

	buildSecrets, err := client.GetBuildSecrets(params, clientBearerToken)
	if err == nil && buildSecrets.Payload != nil && len(buildSecrets.Payload) == 2 {
		if strings.EqualFold(*buildSecrets.Payload[0].Name, build.Secret1) &&
			strings.EqualFold(buildSecrets.Payload[0].Status, expectedStatus) &&
			strings.EqualFold(*buildSecrets.Payload[1].Name, build.Secret2) &&
			strings.EqualFold(buildSecrets.Payload[1].Status, expectedStatus) {
			return nil
		}
	}

	logger.Info("Build secrets are not listed yet")
	return fmt.Errorf("failed buildSecretsAreListedWithStatus expected %s", expectedStatus)
}

func setSecret(cfg config.Config, secretName, secretValue string) error {
	logger.Debugf("setSecret %s with value %s", secretName, secretValue)
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	secretParameters := models.SecretParameters{
		SecretValue: &secretValue,
	}

	params := applicationClient.NewUpdateBuildSecretsSecretValueParams().
		WithAppName(defaults.App2Name).
		WithSecretName(secretName).
		WithSecretValue(&secretParameters).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetApplicationClient(cfg)

	_, err := client.UpdateBuildSecretsSecretValue(params, clientBearerToken)
	if err != nil {
		return fmt.Errorf("failed to set secret %s. Error: %v", secretName, err)
	}

	return nil
}
