package happypath

import (
	apiclient "github.com/equinor/radix-cicd-canary/generated-client/client/platform"
	models "github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
)

func registerApplicationWithNoDeployKey(env env.Env) (bool, error) {
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

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetPlatformClient(env)

	registerApplicationOK, err := client.RegisterApplication(params, clientBearerToken)
	return err == nil && registerApplicationOK.Payload.PublicKey != "", err
}
