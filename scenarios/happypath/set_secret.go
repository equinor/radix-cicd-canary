package happypath

import (
	environmentclient "github.com/equinor/radix-cicd-canary-golang/generated-client/client/environment"
	models "github.com/equinor/radix-cicd-canary-golang/generated-client/models"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary-golang/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

const (
	setSecretTestName = "SetSecret"
)

func setSecret() string {

	test.WaitForCheckFunc(isDeploymentConsistent)

	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := environmentclient.NewChangeEnvironmentComponentSecretParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(app2Name).
		WithEnvName(app2EnvironmentName).
		WithComponentName(app2ComponentName).
		WithSecretName(app2SecretName).
		WithComponentSecret(
			&models.SecretParameters{
				SecretValue: stringPtr(app2SecretValue),
			})

	clientBearerToken := httpUtils.GetClientBearerToken()
	client := httpUtils.GetEnvironmentClient()

	_, err := client.ChangeEnvironmentComponentSecret(params, clientBearerToken)
	if err != nil {
		addTestError(setSecretTestName)
		log.Errorf("Error calling ChangeEnvironmentComponentSecret for application %s: %v", app2Name, err)
	} else {
		addTestSuccess(setSecretTestName)
		log.Infof("Test success for set secret for application %s", app2Name)
	}

	return setSecretTestName
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
