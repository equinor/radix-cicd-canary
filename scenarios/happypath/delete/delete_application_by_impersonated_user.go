package delete

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	commonErrors "github.com/equinor/radix-common/utils/errors"
	log "github.com/sirupsen/logrus"
)

// Applications Tests that we are able to delete applications
func Applications(cfg config.Config, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	var errs []error
	for _, appName := range []string{defaults.App1Name, defaults.App2Name, defaults.App4Name} {
		err := application.DeleteByImpersonatedUser(cfg, appName, logger)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return commonErrors.Concat(errs)
}
