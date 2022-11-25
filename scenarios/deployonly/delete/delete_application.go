package delete

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	log "github.com/sirupsen/logrus"
)

// Applications Tests that we are able to delete applications
func Applications(cfg config.Config, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	return application.DeleteByImpersonatedUser(cfg, defaults.App3Name, logger)
}
