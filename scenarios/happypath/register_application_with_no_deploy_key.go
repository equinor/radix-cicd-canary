package happypath

import (
	apiclient "github.com/equinor/radix-cicd-canary-golang/generated-client/client/platform"
	models "github.com/equinor/radix-cicd-canary-golang/generated-client/models"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary-golang/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

func registerApplicationWithNoDeployKey() string {
	const (
		testName = "RegisterApplicationWithNoDeployKey"
		basePath = "/api/v1"
	)

	log.Infof("Starting RegisterApplication with no deploy key...")

	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	appName := app1Name
	appRepo := app1Repository
	appSharedSecret := app1SharedSecret

	bodyParameters := models.ApplicationRegistration{
		Name:         &appName,
		Repository:   &appRepo,
		SharedSecret: &appSharedSecret,
	}

	params := apiclient.NewRegisterApplicationParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithApplicationRegistration(&bodyParameters)

	clientBearerToken := httpUtils.GetClientBearerToken()
	client := httpUtils.GetPlatformClient()

	registerApplicationOK, err := client.RegisterApplication(params, clientBearerToken)
	if err != nil {
		addTestError(testName)
		log.Errorf("Error calling RegisterApplication with no deploy key: %v", err)
	} else {
		if registerApplicationOK.Payload.PublicKey != "" {
			addTestSuccess(testName)
			log.Info("Test success")
		} else {
			addTestError(testName)
			log.Errorf("Error response: public key is empty")
		}
	}

	return testName
}
