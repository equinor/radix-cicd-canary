package teardown

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	log "github.com/sirupsen/logrus"
)

// TearDown Deletes applications
func TearDown(env env.Env, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	for _, appName := range []string{config.App1Name, config.App2Name, config.App4Name} {
		err := application.DeleteByServiceAccount(env, appName, logger)
		if err != nil {
			logger.Debug(err)
		}
	}
	return nil
}
