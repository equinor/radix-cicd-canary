package delete

import (
	"time"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Applications Tests that we are able to delete applications
func Applications(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// HACK Try sleeping for 60 seconds to hopefully preventing kubed from crashing
	logger.Info("Sleeping for 60 seconds before deleteing application to prevent kubed from crashing")
	time.Sleep(60 * time.Second)

	return application.Delete(env, config.App3Name)
}
