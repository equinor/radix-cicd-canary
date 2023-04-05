package register

import (
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ApplicationWithNoDeployKey Tests that we are able to register application
// with no deploy key and that deploy key is generated
func ApplicationWithNoDeployKey(cfg config.Config, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	appName := defaults.App1Name
	appRepo := defaults.App1Repository
	appSharedSecret := defaults.App1SharedSecret
	appCreator := defaults.App1Creator
	appConfigBranch := defaults.App1ConfigBranch
	appConfigurationItem := defaults.App1ConfigurationItem

	err := application.DeleteIfExist(cfg, appName, logger)
	if err != nil {
		return err
	}

	registerApplicationOK, err := application.Register(cfg, appName, appRepo, appSharedSecret, appCreator, "", "", appConfigBranch, appConfigurationItem, []string{cfg.GetImpersonateGroup()})
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("failed to register application %s", appName))
	}

	if registerApplicationOK.Payload.ApplicationRegistration.PublicKey == "" {
		return fmt.Errorf("the Public Key of the registered application %s is empty", appName)
	}

	return test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return application.IsDefined(cfg, appName)
	}, logger)
}
