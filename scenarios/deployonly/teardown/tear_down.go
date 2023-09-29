package teardown

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/rs/zerolog/log"
)

// TearDown Deletes applications
func TearDown(cfg config.Config, suiteName string) error {
	logger := log.With().Str("suite", suiteName).Logger() //WithFields(log.Fields{"Suite": suiteName})
	err := application.DeleteByServiceAccount(cfg, defaults.App3Name, logger)
	if err != nil {
		logger.Debug().Err(err).Msg("Teardown failure")
	}
	return nil
}
