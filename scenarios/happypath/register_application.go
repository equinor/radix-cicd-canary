package happypath

import (
	applicationclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	apiclient "github.com/equinor/radix-cicd-canary/generated-client/client/platform"
	models "github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

func registerApplication(env env.Env) (bool, error) {
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

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetPlatformClient(env)

	_, err := client.RegisterApplication(params, clientBearerToken)
	if err != nil {
		return false, err
	}

	test.WaitForCheckFuncWithArguments(env, isApplicationDefined, []string{app2Name})
	return true, nil
}

func isApplicationDefined(env env.Env, args []string) (bool, interface{}) {
	appName := args[0]
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := applicationclient.NewGetApplicationParams().
		WithAppName(appName).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	_, err := client.GetApplication(params, clientBearerToken)
	if err == nil {
		return true, nil
	}

	log.Info("Application is not defined")
	return false, nil
}
