package alias

import (
	"errors"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/http"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
)

// DefaultResponding Checks if default alias of application is responding
func DefaultResponding(env envUtil.Env, suiteName string) error {
	ok, publicDomainName := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return application.TryGetPublicDomainName(env, config.App3Name, config.App3EnvironmentName, config.App3Component1Name)
	})

	if !ok {
		return errors.New("public domain name of alias is empty")
	}

	ok, canonicalDomainName := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return application.TryGetCanonicalDomainName(env, config.App3Name, config.App3EnvironmentName, config.App3Component1Name)
	})

	if !ok {
		return errors.New("canonical domain name of alias is empty")
	}

	if application.IsRunningInActiveCluster(publicDomainName.(string), canonicalDomainName.(string)) {
		ok, _ := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
			return application.IsAliasDefined(env, config.App3Name)
		})

		if !ok {
			return errors.New("public alias is not defined")
		}
	}

	ok, _ = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		schema := "https"
		return application.AreResponding(env, http.GetUrl(schema, canonicalDomainName.(string)), http.GetUrl(schema, publicDomainName.(string)))
	})
	return nil
}
