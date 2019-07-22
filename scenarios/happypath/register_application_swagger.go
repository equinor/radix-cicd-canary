package happypath

import (
	apiclient "github.com/equinor/radix-cicd-canary-golang/generated/client/platform"
	models "github.com/equinor/radix-cicd-canary-golang/generated/models"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

func registerApplicationSwagger() string {
	const (
		testName = "RegisterApplicationSwagger"
		basePath = "/api/v1"
	)

	log.Infof("Starting RegisterApplication...")

	radixAPIURL := utils.GetRadixAPIURL()
	impersonateUser := utils.GetImpersonateUser()
	impersonateGroup := utils.GetImpersonateGroup()
	bearerToken := utils.GetBearerToken()

	appName := app2Name
	appRepo := app2Repository
	appSharedSecret := app2SharedSecret

	bodyParameters := models.ApplicationRegistration{
		Name:         &appName,
		Repository:   &appRepo,
		SharedSecret: &appSharedSecret,
		AdGroups:     nil,
		PublicKey:    utils.GetPublicKey(),
		PrivateKey:   utils.GetPrivateKeyBase64(),
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
