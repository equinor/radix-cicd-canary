package alias

import (
	"context"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
)

// DefaultResponding Checks if default alias of application is responding
func DefaultResponding(ctx context.Context, cfg config.Config) error {

	publicDomainName, err := test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (string, error) {
		return application.TryGetPublicDomainName(ctx, cfg, defaults.App3Name, defaults.App3EnvironmentName, defaults.App3Component1Name)
	})
	if err != nil {
		return err
	}

	canonicalDomainName, err := test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (string, error) {
		return application.TryGetCanonicalDomainName(ctx, cfg, defaults.App3Name, defaults.App3EnvironmentName, defaults.App3Component1Name)
	})
	if err != nil {
		return err
	}

	if application.IsRunningInActiveCluster(publicDomainName, canonicalDomainName) {
		err := test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
			return application.IsAliasDefined(ctx, cfg, defaults.App3Name)
		})
		if err != nil {
			return err
		}
	}

	err = test.WaitForCheckFuncOrTimeout(ctx, cfg, func(_ config.Config, ctx context.Context) error {
		schema := "https"
		return application.AreResponding(ctx, http.GetUrl(schema, canonicalDomainName), http.GetUrl(schema, publicDomainName))
	})

	return err
}
