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

	appName := config.App3Name
	appRepo := config.App3Repository
	appSharedSecret := config.App3SharedSecret
	appCreator := config.App3Creator
	appOwner := config.App3Owner

	bodyParameters := models.ApplicationRegistration{
		Name:         &appName,
		Repository:   &appRepo,
		SharedSecret: &appSharedSecret,
		Owner:        &appOwner,
		Creator:      &appCreator,
		AdGroups:     nil,
		PublicKey:    env.GetPublicKeyCanary3(),
		PrivateKey:   env.GetPrivateKeyCanary3(),
	}

	params := apiclient.NewRegisterApplicationParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithApplicationRegistration(&bodyParameters)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetPlatformClient(env)

	_, err := client.RegisterApplication(params, clientBearerToken)
	if err != nil {
		logger.Errorf("%v", err)
		return false, err
	}

	test.WaitForCheckFuncWithArguments(env, isApplicationDefined, []string{config.App3Name})
	return true, nil
}

func isApplicationDefined(env env.Env, args []string) (bool, interface{}) {
	appName := args[0]
	_, err := GetApplication(env, appName)
	if err == nil {
		return true, nil
	}

	logger.Info("Application is not defined")
	return false, nil
}

// GetApplication gets an application by appName
func GetApplication(env env.Env, appName string) (*models.Application, error) {
	params := applicationclient.NewGetApplicationParams().
		WithAppName(appName).
		WithImpersonateUser(env.GetImpersonateUserPointer()).
		WithImpersonateGroup(env.GetImpersonateGroupPointer())
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	result, err := client.GetApplication(params, clientBearerToken)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}
