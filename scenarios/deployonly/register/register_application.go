package register

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Application Tests that we are able to register application
// with deploy key set
func Application(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	appName := config.App3Name
	appRepo := config.App3Repository
	appSharedSecret := config.App3SharedSecret
	appCreator := config.App3Creator
	appOwner := config.App3Owner

	_, err := application.Register(env, appName, appRepo, appSharedSecret, appCreator, appOwner, env.GetPublicKeyCanary3(), env.GetPrivateKeyCanary3())
	if err != nil {
		logger.Errorf("%v", err)
		return false, err
	}

	test.WaitForCheckFuncWithArguments(env, application.IsDefined, []string{config.App3Name})
	return true, nil
}
