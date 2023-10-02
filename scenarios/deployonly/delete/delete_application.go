package delete

import (
	"context"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/rs/zerolog/log"
)

// Applications Tests that we are able to delete applications
func Applications(ctx context.Context, cfg config.Config, suiteName string) error {
	appCtx := log.Ctx(ctx).With().Str("app", defaults.App3Name).Logger().WithContext(ctx)
	return application.DeleteByImpersonatedUser(cfg, defaults.App3Name, appCtx)
}
