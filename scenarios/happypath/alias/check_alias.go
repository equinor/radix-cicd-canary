package alias

import (
	"context"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/rs/zerolog/log"
)

// DefaultResponding Checks if default alias of application is responding
func DefaultResponding(ctx context.Context, cfg config.Config) error {
	appName := defaults.App2Name
	appCtx := log.Ctx(ctx).With().Str("app", appName).Logger().WithContext(ctx)

	publicDomainName, err := test.WaitForCheckFuncWithValueOrTimeout(appCtx, cfg, func(cfg config.Config, ctx context.Context) (string, error) {
		return application.TryGetPublicDomainName(ctx, cfg, appName, defaults.App2EnvironmentName, defaults.App2Component1Name)
	})

	if err != nil {
		return err
	}

	canonicalDomainName, err := test.WaitForCheckFuncWithValueOrTimeout(appCtx, cfg, func(cfg config.Config, ctx context.Context) (string, error) {
		return application.TryGetCanonicalDomainName(ctx, cfg, appName, defaults.App2EnvironmentName, defaults.App2Component1Name)
	})

	if err != nil {
		return err
	}

	if application.IsRunningInActiveCluster(publicDomainName, canonicalDomainName) {
		err := test.WaitForCheckFuncOrTimeout(appCtx, cfg, func(cfg config.Config, ctx context.Context) error {
			return application.IsAliasDefined(ctx, cfg, appName)
		})

		if err != nil {
			return err
		}
	}

	return test.WaitForCheckFuncOrTimeout(appCtx, cfg, func(_ config.Config, ctx context.Context) error {
		schema := "https"
		return application.AreResponding(ctx, http.GetUrl(schema, canonicalDomainName), http.GetUrl(schema, publicDomainName))
	})
}
