package delete

import (
	"context"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	commonErrors "github.com/equinor/radix-common/utils/errors"
	"github.com/rs/zerolog/log"
)

// Applications Tests that we are able to delete applications
func Applications(ctx context.Context, cfg config.Config, suiteName string) error {
	var errs []error
	for _, appName := range []string{defaults.App1Name, defaults.App2Name, defaults.App4Name} {
		appCtx := log.Ctx(ctx).With().Str("app", appName).Logger().WithContext(ctx)
		err := application.DeleteByImpersonatedUser(appCtx, cfg, appName)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return commonErrors.Concat(errs)
}
