package register

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Application Tests that we are able to register application
// with deploy key set
func Application(env envUtil.Env, suiteName string) error {
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
		return err
	}

	return test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) error {
		return application.IsDefined(env, config.App2Name)
	})
}
