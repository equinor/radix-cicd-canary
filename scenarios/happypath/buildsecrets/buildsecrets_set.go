package buildsecrets

import (
	"strings"

	applicationclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	"github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/build"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Set Tests that we are able to successfully set build secrets
func Set(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Trigger build to apply RA with build secrets
	ok := httpUtils.TriggerWebhookPush(env, config.App2BranchToBuildFrom, config.App2CommitID, config.App2SSHRepository, config.App2SharedSecret)
	if !ok {
		return false, nil
	}

	logger.Info("Job was triggered to apply RA")

	// Get job
	ok, jobSummary := test.WaitForCheckFuncWithArguments(env, build.IsJobListedWithStatus, []string{"Failed"})
	if !ok {
		return false, nil
	}

	jobName := (jobSummary.(*models.JobSummary)).Name
	job := build.Job(env, jobName)
	if len(job.Steps) != 2 {
		logger.Error("Job should not contain any build step")
		return false, nil
	}

	// First job failed, due to missing build secrets, as expected in test
	// Set build secrets
	ok, _ = test.WaitForCheckFuncWithArguments(env, buildSecretsAreListedWithStatus, []string{"Pending"})
	if !ok {
		return false, nil
	}

	ok, err := setSecret(env, build.Secret1, build.Secret1Value)
	if !ok {
		return false, err
	}

	ok, err = setSecret(env, build.Secret2, build.Secret2Value)
	if !ok {
		return false, err
	}

	ok, _ = test.WaitForCheckFuncWithArguments(env, buildSecretsAreListedWithStatus, []string{"Consistent"})
	if !ok {
		return false, nil
	}

	return true, nil
}

func buildSecretsAreListedWithStatus(env env.Env, args []string) (bool, interface{}) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	expectedStatus := args[0]
	params := applicationclient.NewGetBuildSecretsParams().
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
			return true, nil
		}
	}

	logger.Info("Build secrets are not listed yet")
	return false, nil
}

func setSecret(env env.Env, secretName, secretValue string) (bool, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	secretParameters := models.SecretParameters{
		SecretValue: &secretValue,
	}

	params := applicationclient.NewUpdateBuildSecretsSecretValueParams().
		WithAppName(config.App2Name).
		WithSecretName(secretName).
		WithSecretValue(&secretParameters).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	_, err := client.UpdateBuildSecretsSecretValue(params, clientBearerToken)
	if err != nil {
		return false, err
	}

	return true, nil
}
