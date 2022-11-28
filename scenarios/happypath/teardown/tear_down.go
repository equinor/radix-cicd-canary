package teardown

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	log "github.com/sirupsen/logrus"
)

// TearDown Deletes applications
func TearDown(cfg config.Config, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	for _, appName := range []string{defaults.App1Name, defaults.App2Name, defaults.App4Name} {
		err := application.DeleteByServiceAccount(cfg, appName, logger)
		if err != nil {
			logger.Debug(err)
		}
	}
	return nil
}
