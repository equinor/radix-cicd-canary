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

// ApplicationWithNoDeployKey Tests that we are able to register application
// with no deploy key and that deploy key is generated
func ApplicationWithNoDeployKey(env envUtil.Env, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	appName := config.App1Name
	appRepo := config.App1Repository
	appSharedSecret := config.App1SharedSecret
	appCreator := config.App1Creator
	appConfigBranch := config.App1ConfigBranch
	appConfigurationItem := config.App1ConfigurationItem

	err := application.DeleteIfExist(env, appName, logger)
	if err != nil {
		return err
	}

	registerApplicationOK, err := application.Register(env, appName, appRepo, appSharedSecret, appCreator, "", "", appConfigBranch, appConfigurationItem)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("failed to register application %s", appName))
	}

	if registerApplicationOK.Payload.ApplicationRegistration.PublicKey == "" {
		return fmt.Errorf("the Public Key of the registered application %s is empty", appName)
	}

	return test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) error {
		return application.IsDefined(env, appName)
	}, logger)
}
