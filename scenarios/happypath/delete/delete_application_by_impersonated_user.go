package delete

import (
	"context"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	commonErrors "github.com/equinor/radix-common/utils/errors"
)

// Applications Tests that we are able to delete applications
func Applications(ctx context.Context, cfg config.Config) error {
	var errs []error
	for _, appName := range []string{defaults.App1Name, defaults.App2Name, defaults.App4Name} {
		err := application.DeleteByImpersonatedUser(ctx, cfg, appName)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return commonErrors.Concat(errs)
}
