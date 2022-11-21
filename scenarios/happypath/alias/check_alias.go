package alias

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

// DefaultResponding Checks if default alias of application is responding
func DefaultResponding(env envUtil.Env, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	publicDomainName, err := test.WaitForCheckFuncWithValueOrTimeout(env, func(env envUtil.Env) (string, error) {
		return application.TryGetPublicDomainName(env, config.App2Name, config.App2EnvironmentName, config.App2Component1Name)
	}, logger)

	if err != nil {
		return err
	}

	canonicalDomainName, err := test.WaitForCheckFuncWithValueOrTimeout(env, func(env envUtil.Env) (string, error) {
		return application.TryGetCanonicalDomainName(env, config.App2Name, config.App2EnvironmentName, config.App2Component1Name)
	}, logger)

	if err != nil {
		return err
	}

	if application.IsRunningInActiveCluster(publicDomainName, canonicalDomainName) {
		err := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) error {
			return application.IsAliasDefined(env, config.App2Name)
		}, logger)

		if err != nil {
			return err
		}
	}

	return test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) error {
		schema := "https"
		return application.AreResponding(env, http.GetUrl(schema, canonicalDomainName), http.GetUrl(schema, publicDomainName))
	}, logger)
}
