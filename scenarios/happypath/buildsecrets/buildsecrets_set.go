package buildsecrets

import (
	"fmt"
	"strings"

	applicationClient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	"github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/build"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Set Tests that we are able to successfully set build secrets
func Set(env envUtil.Env, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Trigger build to apply RA with build secrets
	err := httpUtils.TriggerWebhookPush(env, config.App2BranchToBuildFrom, config.App2CommitID, config.App2SSHRepository, config.App2SharedSecret)
	if err != nil {
		return err
	}

	logger.Info("Job was triggered to apply RA")

	// Get job
	jobSummary, err := test.WaitForCheckFuncWithValueOrTimeout(env, func(env envUtil.Env) (*models.JobSummary, error) {
		return job.IsListedWithStatus(env, config.App2Name, "Failed", logger)
	}, logger)
	if err != nil {
		return err
	}

	jobName := jobSummary.Name
	job, _ := job.Get(env, config.App2Name, jobName)

	expectedSteps := []string{
		"clone-config",
		"prepare-pipelines",
		"radix-pipeline"}

	if len(job.Steps) != len(expectedSteps) {
		return fmt.Errorf("job should not contain any build step")
	}

	// First job failed, due to missing build secrets, as expected in test
	// Set build secrets
	err = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) error {
		return buildSecretsAreListedWithStatus(env, "Pending")
	}, logger)

	if err != nil {
		return err
	}

	err = setSecret(env, build.Secret1, build.Secret1Value)
	if err != nil {
		return err
	}

	err = setSecret(env, build.Secret2, build.Secret2Value)
	if err != nil {
		return err
	}

	return test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) error {
		return buildSecretsAreListedWithStatus(env, "Consistent")
	}, logger)
}

func buildSecretsAreListedWithStatus(env envUtil.Env, expectedStatus string) error {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := applicationClient.NewGetBuildSecretsParams().
		WithAppName(config.App2Name).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	buildsecrets, err := client.GetBuildSecrets(params, clientBearerToken)
	if err == nil && buildsecrets.Payload != nil && len(buildsecrets.Payload) == 2 {
		if strings.EqualFold(*buildsecrets.Payload[0].Name, build.Secret1) &&
			strings.EqualFold(buildsecrets.Payload[0].Status, expectedStatus) &&
			strings.EqualFold(*buildsecrets.Payload[1].Name, build.Secret2) &&
			strings.EqualFold(buildsecrets.Payload[1].Status, expectedStatus) {
			return nil
		}
	}

	logger.Info("Build secrets are not listed yet")
	return fmt.Errorf("failed buildSecretsAreListedWithStatus expected %s", expectedStatus)
}

func setSecret(env envUtil.Env, secretName, secretValue string) error {
	log.Debugf("setSecret %s with value %s", secretName, secretValue)
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	secretParameters := models.SecretParameters{
		SecretValue: &secretValue,
	}

	params := applicationClient.NewUpdateBuildSecretsSecretValueParams().
		WithAppName(config.App2Name).
		WithSecretName(secretName).
		WithSecretValue(&secretParameters).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	_, err := client.UpdateBuildSecretsSecretValue(params, clientBearerToken)
	if err != nil {
		return fmt.Errorf("failed to set secret %s. Error: %v", secretName, err)
	}

	return nil
}
