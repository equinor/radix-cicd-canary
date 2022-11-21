package alias

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
)

// DefaultResponding Checks if default alias of application is responding
func DefaultResponding(env envUtil.Env, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	publicDomainName, err := test.WaitForCheckFuncWithValueOrTimeout(env, func(env envUtil.Env) (string, error) {
		return application.TryGetPublicDomainName(env, config.App3Name, config.App3EnvironmentName, config.App3Component1Name)
	}, logger)
	if err != nil {
		return err
	}

	canonicalDomainName, err := test.WaitForCheckFuncWithValueOrTimeout(env, func(env envUtil.Env) (string, error) {
		return application.TryGetCanonicalDomainName(env, config.App3Name, config.App3EnvironmentName, config.App3Component1Name)
	}, logger)
	if err != nil {
		return err
	}

	if application.IsRunningInActiveCluster(publicDomainName, canonicalDomainName) {
		err := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) error {
			return application.IsAliasDefined(env, config.App3Name, logger)
		}, logger)
		if err != nil {
			return err
		}
	}

	err = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) error {
		schema := "https"
		return application.AreResponding(env, logger, http.GetUrl(schema, canonicalDomainName), http.GetUrl(schema, publicDomainName))
	}, logger)

	return err
}
