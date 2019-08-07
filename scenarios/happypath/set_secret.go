package happypath

import (
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/client/environment"
	models "github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

func setSecret(env env.Env) (bool, error) {
	test.WaitForCheckFunc(env, isDeploymentConsistent)

	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := environmentclient.NewChangeEnvironmentComponentSecretParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(app2Name).
		WithEnvName(app2EnvironmentName).
		WithComponentName(app2Component2Name).
		WithSecretName(app2SecretName).
		WithComponentSecret(
			&models.SecretParameters{
				SecretValue: stringPtr(app2SecretValue),
			})

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetEnvironmentClient(env)

	_, err := client.ChangeEnvironmentComponentSecret(params, clientBearerToken)
	if err != nil {
		log.Errorf("Error calling ChangeEnvironmentComponentSecret for application %s: %v", app2Name, err)
	}

	return err == nil, err
}

func isDeploymentConsistent(env env.Env, args []string) (bool, interface{}) {
	environmentDetails := getEnvironmentDetails(env)
	if environmentDetails != nil && environmentDetails.ActiveDeployment != nil && environmentDetails.Status != "" {
		log.Info("Deployment is consistent. We can set the secret.")
		return true, nil
	}

	return false, nil
}

func getEnvironmentDetails(env env.Env) *models.Environment {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := environmentclient.NewGetEnvironmentParams().
		WithAppName(app2Name).
		WithEnvName(app2EnvironmentName).
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
