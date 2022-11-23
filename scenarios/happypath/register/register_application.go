package register

import (
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Application Tests that we are able to register application
// with deploy key set
func Application(env envUtil.Env, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	appName := config.App2Name
	appRepo := config.App2Repository
	appSharedSecret := config.App2SharedSecret
	appCreator := config.App2Creator
	appConfigBranch := config.App2ConfigBranch
	appConfigurationItem := config.App2ConfigurationItem

	err := application.DeleteIfExist(env, appName, logger)
	if err != nil {
		return err
	}

	_, err = application.Register(env, appName, appRepo, appSharedSecret, appCreator, env.GetPublicKey(), env.GetPrivateKey(), appConfigBranch, appConfigurationItem)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("failed to register application %s", appName))
	}

	return test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) error {
		return application.IsDefined(env, appName)
	}, logger)
}
