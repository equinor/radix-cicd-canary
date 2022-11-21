package register

import (
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/pkg/errors"
)

// ApplicationWithMainConfigBranch Tests that we are able to register application
// with no deploy key and that deploy key is generated
func ApplicationWithMainConfigBranch(env envUtil.Env, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	appName := config.App4Name
	appRepo := config.App4Repository
	appSharedSecret := config.App4SharedSecret
	appCreator := config.App4Creator
	appConfigBranch := config.App4ConfigBranch
	appConfigurationItem := config.App4ConfigurationItem

	_, err := application.Register(env, appName, appRepo, appSharedSecret, appCreator, env.GetPublicKeyCanary4(), env.GetPrivateKeyCanary4(), appConfigBranch, appConfigurationItem)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("failed to register application %s", appName))
	}

	return test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) error {
		return application.IsDefined(env, appName)
	}, logger)
}
