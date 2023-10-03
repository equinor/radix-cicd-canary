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
	err := application.DeleteByServiceAccount(ctx, cfg, defaults.App3Name)
	if err != nil {
		log.Ctx(ctx).Debug().Str("app", defaults.App3Name).Stack().Err(err).Msg("Teardown failure")
	}
	return nil
}
