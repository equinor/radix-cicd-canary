package register

import (
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ApplicationWithMainConfigBranch Tests that we are able to register application
// with no deploy key and that deploy key is generated
func ApplicationWithMainConfigBranch(env envUtil.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})
	appName := config.App4Name
	appRepo := config.App4Repository
	appSharedSecret := config.App4SharedSecret
	appOwner := config.App4Owner
	appCreator := config.App4Creator
	appWbs := config.App4Wbs
	appConfigBranch := config.App4ConfigBranch

	_, err := application.Register(env, appName, appRepo, appSharedSecret, appCreator, appOwner, env.GetPublicKeyCanary4(), env.GetPrivateKeyCanary4(), appWbs, appConfigBranch)
	if err != nil {
		return false, errors.WithMessage(err, fmt.Sprintf("failed to register application %s", appName))
	}

	ok, _ := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return application.IsDefined(env, appName)
	})

	if !ok {
		return false, fmt.Errorf("application %s is not defined", appName)
	}

	return true, nil
}
