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
		err := application.DeleteByServiceAccount(ctx, cfg, appName)
		if err != nil {
			log.Ctx(ctx).Debug().Str("app", appName).Err(err).Msgf("Teardown failed")
		}
	}
	return nil
}
