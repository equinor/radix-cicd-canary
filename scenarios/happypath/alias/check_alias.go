package alias

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

// DefaultResponding Checks if default alias of application is responding
func DefaultResponding(cfg config.Config, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	publicDomainName, err := test.WaitForCheckFuncWithValueOrTimeout(cfg, func(cfg config.Config) (string, error) {
		return application.TryGetPublicDomainName(cfg, defaults.App2Name, defaults.App2EnvironmentName, defaults.App2Component1Name)
	}, logger)

	if err != nil {
		return err
	}

	canonicalDomainName, err := test.WaitForCheckFuncWithValueOrTimeout(cfg, func(cfg config.Config) (string, error) {
		return application.TryGetCanonicalDomainName(cfg, defaults.App2Name, defaults.App2EnvironmentName, defaults.App2Component1Name)
	}, logger)

	if err != nil {
		return err
	}

	if application.IsRunningInActiveCluster(publicDomainName, canonicalDomainName) {
		err := test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
			return application.IsAliasDefined(cfg, defaults.App2Name, logger)
		}, logger)

		if err != nil {
			return err
		}
	}

	return test.WaitForCheckFuncOrTimeout(cfg, func(_ config.Config) error {
		schema := "https"
		return application.AreResponding(logger, http.GetUrl(schema, canonicalDomainName), http.GetUrl(schema, publicDomainName))
	}, logger)
}
