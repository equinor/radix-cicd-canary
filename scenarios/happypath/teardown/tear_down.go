package teardown

import (
	"context"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/rs/zerolog/log"
)

// TearDown Deletes applications
func TearDown(ctx context.Context, cfg config.Config) error {
	for _, appName := range []string{defaults.App1Name, defaults.App2Name, defaults.App4Name} {
		appCtx := log.Ctx(ctx).With().Str("app", appName).Logger().WithContext(ctx)
		err := application.DeleteByServiceAccount(appCtx, cfg, appName)
		if err != nil {
			log.Ctx(appCtx).Debug().Err(err).Msg("Teardown failed")
		}
	}
	return nil
}
