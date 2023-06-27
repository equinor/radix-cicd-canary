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

// Application Tests that we are able to register application
// with deploy key set
func Application(cfg config.Config, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	appName := defaults.App2Name
	appRepo := defaults.App2Repository
	appSharedSecret := defaults.App2SharedSecret
	appCreator := defaults.App2Creator
	appConfigBranch := defaults.App2ConfigBranch
	appConfigurationItem := defaults.App2ConfigurationItem

	err := application.DeleteIfExist(cfg, appName, logger)
	if err != nil {
		return err
	}

	_, err = application.Register(cfg, appName, appRepo, appSharedSecret, appCreator, appConfigBranch, appConfigurationItem, cfg.GetAppAdminGroup())
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("failed to register application %s", appName))
	}

	err = test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return application.IsDefined(cfg, appName)
	}, logger)
	if err != nil {
		return err
	}

	err = application.RegenerateDeployKey(cfg, appName, cfg.GetPrivateKey(), "", logger)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("failed to regenerate deploy key for application %s", appName))
	}

	return test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return application.HasDeployKey(cfg, appName, cfg.GetPublicKey(), logger)
	}, logger)
}
