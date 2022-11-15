package register

import (
	"errors"
	"fmt"

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
	appCreator := config.App1Creator
	appConfigBranch := config.App1ConfigBranch
	appConfigurationItem := config.App1ConfigurationItem

	registerApplicationOK, err := application.Register(env, appName, appRepo, appSharedSecret, appCreator, "", "", appConfigBranch, appConfigurationItem)
	if err != nil {
		logger.Errorf("%v", err)
		return false, err
	}

	ok, _ := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return application.IsDefined(env, config.App2Name)
	})

	if !ok {
		return false, errors.New(fmt.Sprintf("failed to get details of registered application %s", config.App2Name))
	}

	return registerApplicationOK.Payload.ApplicationRegistration.PublicKey != "", err
}
