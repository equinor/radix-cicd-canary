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

	ok, publicDomainName := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return application.IsPublicDomainNameDefined(env, config.App3Name, config.App3EnvironmentName, config.App3Component1Name)
	})

	if !ok {
		return false, fmt.Errorf("Public domain name of alias is empty")
	}

	ok, canonicalDomainName := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return application.IsCanonicalDomainNameDefined(env, config.App3Name, config.App3EnvironmentName, config.App3Component1Name)
	})

	if !ok {
		return false, fmt.Errorf("Canonical domain name of alias is empty")
	}

	if application.IsRunningInActiveCluster(publicDomainName.(string), canonicalDomainName.(string)) {
		ok, _ := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
			return application.IsAliasDefined(env, config.App3Name)
		})

		if !ok {
			return false, fmt.Errorf("Public alias is not defined")
		}
	}

	ok, _ = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return application.AreResponding(env, canonicalDomainName.(string), publicDomainName.(string))
	})
	return ok, nil
}
