package happypath

import (
	apiclient "github.com/equinor/radix-cicd-canary-golang/generated-client/client/platform"
	models "github.com/equinor/radix-cicd-canary-golang/generated-client/models"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils/env"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

func registerApplication() string {
	const (
		testName = "RegisterApplication"
		basePath = "/api/v1"
	)

	log.Infof("Starting RegisterApplication...")

	radixAPIURL := env.GetRadixAPIURL()
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()
	bearerToken := env.GetBearerToken()

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
	clientBearerToken := httptransport.BearerToken(bearerToken)
	schemes := []string{"https"}

	transport := httptransport.New(radixAPIURL, basePath, schemes)
	client := apiclient.New(transport, strfmt.Default)

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
