package happypath

import (
	apiclient "github.com/equinor/radix-cicd-canary-golang/generated-client/client/platform"
	models "github.com/equinor/radix-cicd-canary-golang/generated-client/models"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary-golang/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

func registerApplication() string {
	const (
		testName = "RegisterApplication"
		basePath = "/api/v1"
	)

	log.Infof("Starting RegisterApplication...")

	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	appName := app2Name
	appRepo := app2Repository
	appSharedSecret := app2SharedSecret

	bodyParameters := models.ApplicationRegistration{
		Name:         &appName,
		Repository:   &appRepo,
		SharedSecret: &appSharedSecret,
		AdGroups:     nil,
		PublicKey:    env.GetPublicKey(),
		PrivateKey:   env.GetPrivateKey(),
	}

	params := apiclient.NewRegisterApplicationParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithApplicationRegistration(&bodyParameters)

	clientBearerToken := httpUtils.GetClientBearerToken()
	client := httpUtils.GetPlatformClient()

	_, err := client.RegisterApplication(params, clientBearerToken)
	if err != nil {
		addTestError(testName)
		log.Errorf("Error calling RegisterApplication: %v", err)
	} else {
		addTestSuccess(testName)
		log.Info("Test success")
	}

	return testName
}
