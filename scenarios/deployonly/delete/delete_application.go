package delete

import (
	"context"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
)

// Applications Tests that we are able to delete applications
func Applications(ctx context.Context, cfg config.Config) error {
	return application.DeleteByImpersonatedUser(ctx, cfg, defaults.App3Name)
}
