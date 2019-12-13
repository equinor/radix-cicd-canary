package secret

import (
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/client/environment"
	models "github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Set Test that we are able to set secret
func Set(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	test.WaitForCheckFuncOrTimeout(env, isDeploymentConsistent)

	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := environmentclient.NewChangeEnvironmentComponentSecretParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(config.App2Name).
		WithEnvName(config.App2EnvironmentName).
		WithComponentName(config.App2Component2Name).
		WithSecretName(config.App2SecretName).
		WithComponentSecret(
			&models.SecretParameters{
				SecretValue: stringPtr(config.App2SecretValue),
			})

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetEnvironmentClient(env)

	_, err := client.ChangeEnvironmentComponentSecret(params, clientBearerToken)
	if err != nil {
		logger.Errorf("Error calling ChangeEnvironmentComponentSecret for application %s: %v", config.App2Name, err)
	}

	return err == nil, err
}

func isDeploymentConsistent(env env.Env) (bool, interface{}) {
	environmentDetails := getEnvironmentDetails(env)
	if environmentDetails != nil &&
		environmentDetails.ActiveDeployment != nil &&
		environmentDetails.Status != "" &&
		len(environmentDetails.Secrets) > 0 {
		logger.Info("Deployment is consistent. We can set the secret.")
		return true, nil
	}

	return false, nil
}

func getEnvironmentDetails(env env.Env) *models.Environment {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := environmentclient.NewGetEnvironmentParams().
		WithAppName(config.App2Name).
		WithEnvName(config.App2EnvironmentName).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetEnvironmentClient(env)

	environmentDetails, err := client.GetEnvironment(params, clientBearerToken)
	if err == nil && environmentDetails.Payload != nil {
		return environmentDetails.Payload
	}

	return nil
}

func stringPtr(str string) *string {
	return &str
}
