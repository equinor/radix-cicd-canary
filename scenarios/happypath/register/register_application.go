package register

import (
	applicationclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	apiclient "github.com/equinor/radix-cicd-canary/generated-client/client/platform"
	models "github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Application Tests that we are able to register application
// with deploy key set
func Application(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	appName := config.App2Name
	appRepo := config.App2Repository
	appSharedSecret := config.App2SharedSecret

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

	test.WaitForCheckFuncWithArguments(env, isApplicationDefined, []string{config.App2Name})
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

	logger.Info("Application is not defined")
	return false, nil
}
