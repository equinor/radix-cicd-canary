package register

import (
	"errors"
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Application Tests that we are able to register application
// with deploy key set
func Application(env envUtil.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	appName := config.App2Name
	appRepo := config.App2Repository
	appSharedSecret := config.App2SharedSecret
	appCreator := config.App2Creator
	appConfigBranch := config.App2ConfigBranch
	appConfigurationItem := config.App2ConfigurationItem

	_, err := application.Register(env, appName, appRepo, appSharedSecret, appCreator, env.GetPublicKey(), env.GetPrivateKey(), appConfigBranch, appConfigurationItem)
	if err != nil {
		logger.Errorf("%v", err)
		return false, err
	}

	ok, _ := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return application.IsDefined(env, config.App2Name)
	})

	if !ok {
		return false, errors.New(fmt.Sprintf("failed to get update details on registered application %s", appName))
	}

	return true, nil
}
