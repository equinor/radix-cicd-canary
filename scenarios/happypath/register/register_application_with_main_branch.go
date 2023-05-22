package register

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/pkg/errors"
)

// ApplicationWithMainConfigBranch Tests that we are able to register the application
func ApplicationWithMainConfigBranch(cfg config.Config, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	appName := defaults.App4Name
	appRepo := defaults.App4Repository
	appSharedSecret := defaults.App4SharedSecret
	appCreator := defaults.App4Creator
	appConfigBranch := defaults.App4ConfigBranch
	appConfigurationItem := defaults.App4ConfigurationItem

	err := application.DeleteIfExist(cfg, appName, logger)
	if err != nil {
		return err
	}

	_, err = application.Register(cfg, appName, appRepo, appSharedSecret, appCreator, appConfigBranch, appConfigurationItem, []string{cfg.GetImpersonateGroup()})
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("failed to register application %s", appName))
	}

	err = test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return application.IsDefined(cfg, appName)
	}, logger)
	if err != nil {
		return err
	}

	err = application.RegenerateDeployKey(cfg, appName, cfg.GetPrivateKeyCanary4(), "", logger)
	if err != nil {
		return err
	}

	return test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return application.HasDeployKey(cfg, appName, cfg.GetPublicKeyCanary4(), logger)
	}, logger)
}
