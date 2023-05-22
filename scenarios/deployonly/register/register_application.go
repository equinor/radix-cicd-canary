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

	_, err = application.Register(cfg, appName, appRepo, appSharedSecret, appCreator, appConfigBranch, appConfigurationItem, []string{cfg.GetImpersonateGroup()})
	if err != nil {
		return err
	}

	return test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return application.RegenerateDeployKey(cfg, appName, cfg.GetPrivateKeyCanary3(), "some-secret", logger)
	}, logger)
}
