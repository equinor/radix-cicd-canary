package alias

import (
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// DefaultResponding Checks if default alias of application is responding
func DefaultResponding(env envUtil.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	publicDomainName := application.GetPublicDomainName(env, config.App3Name, config.App3EnvironmentName, config.App3Component1Name)
	if publicDomainName == "" {
		return false, fmt.Errorf("Public domain name of alias is empty")
	}

	canonicalDomainName := application.GetPublicDomainName(env, config.App3Name, config.App3EnvironmentName, config.App3Component1Name)
	if publicDomainName == "" {
		return false, fmt.Errorf("Public domain name of alias is empty")
	}

	if application.IsRunningInActiveCluster(publicDomainName, canonicalDomainName) {
		ok, _ := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
			return application.IsAliasDefined(env, config.App3Name)
		})

		if !ok {
			return false, fmt.Errorf("Public alias is not defined")
		}
	}

	ok, _ := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return application.AreResponding(env, config.App3Name, canonicalDomainName, publicDomainName)
	})
	return ok, nil
}
