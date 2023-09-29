package teardown

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/rs/zerolog/log"
)

// TearDown Deletes applications
func TearDown(cfg config.Config, suiteName string) error {
	logger := log.With().Str("suite", suiteName).Logger()
	for _, appName := range []string{defaults.App1Name, defaults.App2Name, defaults.App4Name} {
		err := application.DeleteByServiceAccount(cfg, appName, logger)
		if err != nil {
			logger.Debug().Err(err).Msg("Teardown failed")
		}
	}
	return nil
}
