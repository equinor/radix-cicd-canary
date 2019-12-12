package register

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
)

// ApplicationWithNoDeployKey Tests that we are able to register application
// with no deploy key and that deploy key is generated
func ApplicationWithNoDeployKey(env env.Env, suiteName string) (bool, error) {
	appName := config.App1Name
	appRepo := config.App1Repository
	appSharedSecret := config.App1SharedSecret
	appOwner := config.App1Owner
	appCreator := config.App1Creator

	registerApplicationOK, err := application.Register(env, appName, appRepo, appSharedSecret, appCreator, appOwner, "", "")
	if err != nil {
		logger.Errorf("%v", err)
		return false, err
	}

	test.WaitForCheckFuncWithArguments(env, application.IsDefined, []string{config.App1Name})
	return registerApplicationOK.Payload.PublicKey != "", err
}
