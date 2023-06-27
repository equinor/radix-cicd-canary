package register

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

// Application Tests that we are able to register application
// with deploy key set
func Application(cfg config.Config, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	appName := defaults.App3Name
	appRepo := defaults.App3Repository
	appSharedSecret := defaults.App3SharedSecret
	appCreator := defaults.App3Creator
	appConfigurationItem := defaults.App3ConfigurationItem
	appConfigBranch := defaults.App3ConfigBranch

	err := application.DeleteIfExist(cfg, appName, logger)
	if err != nil {
		return err
	}

	_, err = application.Register(cfg, appName, appRepo, appSharedSecret, appCreator, appConfigBranch, appConfigurationItem, cfg.GetAppAdminGroup())
	if err != nil {
		return err
	}

	err = test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return application.IsDefined(cfg, defaults.App3Name)
	}, logger)
	if err != nil {
		return err
	}

	err = application.RegenerateDeployKey(cfg, appName, cfg.GetPrivateKeyCanary3(), "", logger)
	if err != nil {
		return err
	}

	return test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return application.HasDeployKey(cfg, appName, cfg.GetPublicKeyCanary3(), logger)
	}, logger)
}
