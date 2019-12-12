package privateimagehub

import (
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/privateimagehub"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Set runs tests related to private image hub. Expect canary2 to be built and deployed before test run
func Set(env envUtil.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	err := privateimagehub.SetPassword(env, config.App3Name)
	if err != nil {
		return false, fmt.Errorf("Failed to set private image hub password. %v", err)
	}

	return true, nil
}
