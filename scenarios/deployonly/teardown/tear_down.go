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
	appCtx := log.Ctx(ctx).With().Str("app", defaults.App3Name).Logger().WithContext(ctx)
	err := application.DeleteByServiceAccount(appCtx, cfg, defaults.App3Name)
	if err != nil {
		log.Ctx(appCtx).Debug().Err(err).Msg("Teardown failure")
	}
	return nil
}
