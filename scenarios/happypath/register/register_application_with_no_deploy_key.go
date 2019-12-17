package register

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
)

// ApplicationWithNoDeployKey Tests that we are able to register application
// with no deploy key and that deploy key is generated
func ApplicationWithNoDeployKey(env envUtil.Env, suiteName string) (bool, error) {
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

	ok, _ := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return application.IsDefined(env, config.App2Name)
	})

	if !ok {
		return false, nil
	}

	return registerApplicationOK.Payload.PublicKey != "", err
}
