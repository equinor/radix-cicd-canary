package happypath

import (
	environmentclient "github.com/equinor/radix-cicd-canary-golang/generated-client/client/environment"
	models "github.com/equinor/radix-cicd-canary-golang/generated-client/models"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary-golang/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

func setSecret() (bool, error) {
	test.WaitForCheckFunc(isDeploymentConsistent)

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

	clientBearerToken := httpUtils.GetClientBearerToken()
	client := httpUtils.GetEnvironmentClient()

	_, err := client.ChangeEnvironmentComponentSecret(params, clientBearerToken)
	if err != nil {
		log.Errorf("Error calling ChangeEnvironmentComponentSecret for application %s: %v", app2Name, err)
	}

	return err == nil, err
}

func isDeploymentConsistent(args []string) (bool, interface{}) {
	environmentDetails := getEnvironmentDetails()
	if environmentDetails != nil && environmentDetails.ActiveDeployment != nil && environmentDetails.Status != "" {
		log.Info("Deployment is consistent. We can set the secret.")
		return true, nil
	}

	return false, nil
}

func getEnvironmentDetails() *models.Environment {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := environmentclient.NewGetEnvironmentParams().
		WithAppName(app2Name).
		WithEnvName(app2EnvironmentName).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken()
	client := httpUtils.GetEnvironmentClient()

	environmentDetails, err := client.GetEnvironment(params, clientBearerToken)
	if err == nil && environmentDetails.Payload != nil {
		return environmentDetails.Payload
	}

	return nil
}

func stringPtr(str string) *string {
	return &str
}
