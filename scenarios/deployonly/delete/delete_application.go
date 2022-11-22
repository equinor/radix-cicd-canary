package delete

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	log "github.com/sirupsen/logrus"
)

// Applications Tests that we are able to delete applications
func Applications(env env.Env, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	return application.DeleteByImpersonatedUser(env, config.App3Name, logger)
}
