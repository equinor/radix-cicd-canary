package happypath

import (
	apiclient "github.com/equinor/radix-cicd-canary/generated-client/client/platform"
	models "github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
)

func registerApplicationWithNoDeployKey(env env.Env) (bool, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	appName := config.App1Name
	appRepo := config.App1Repository
	appSharedSecret := config.App1SharedSecret

	bodyParameters := models.ApplicationRegistration{
		Name:         &appName,
		Repository:   &appRepo,
		SharedSecret: &appSharedSecret,
	}

	params := apiclient.NewRegisterApplicationParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithApplicationRegistration(&bodyParameters)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetPlatformClient(env)

	registerApplicationOK, err := client.RegisterApplication(params, clientBearerToken)
	if err != nil {
		return false, err
	}

	test.WaitForCheckFuncWithArguments(env, isApplicationDefined, []string{config.App1Name})
	return registerApplicationOK.Payload.PublicKey != "", err
}
