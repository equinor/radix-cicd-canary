package register

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

// Application Tests that we are able to register application
// with deploy key set
func Application(env envUtil.Env, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	appName := config.App3Name
	appRepo := config.App3Repository
	appSharedSecret := config.App3SharedSecret
	appCreator := config.App3Creator
	appConfigurationItem := config.App3ConfigurationItem
	appConfigBranch := config.App3ConfigBranch

	_, err := application.Register(env, appName, appRepo, appSharedSecret, appCreator, env.GetPublicKeyCanary3(), env.GetPrivateKeyCanary3(), appConfigBranch, appConfigurationItem)
	if err != nil {
		return err
	}

	return test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) error {
		return application.IsDefined(env, config.App3Name)
	}, logger)
}
