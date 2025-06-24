package register

import (
	"context"
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/pkg/errors"
)

// Application Tests that we are able to register application
// with deploy key set
func Application(ctx context.Context, cfg config.Config) error {
	appName := defaults.App2Name
	appRepo := defaults.App2Repository
	appSharedSecret := defaults.App2SharedSecret
	appCreator := defaults.App2Creator
	appConfigBranch := defaults.App2ConfigBranch
	appConfigurationItem := defaults.App2ConfigurationItem

	if err := application.DeleteIfExist(ctx, cfg, appName); err != nil {
		return err
	}

	if _, err := application.Register(ctx, cfg, appName, appRepo, appSharedSecret, appCreator, appConfigBranch, appConfigurationItem, cfg.GetAppAdminGroup(), []string{cfg.GetAppReaderGroup()}); err != nil {
		return errors.Wrapf(err, "failed to register application %s", appName)
	}

	if err := test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		return application.IsDefined(ctx, cfg, appName)
	}); err != nil {
		return err
	}

	if err := application.RegenerateDeployKey(ctx, cfg, appName, cfg.GetPrivateKey()); err != nil {
		return fmt.Errorf("failed to regenerated deploy key for application %s: %w", appName, err)
	}

	return test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		return application.HasDeployKey(ctx, cfg, appName, cfg.GetPublicKey())
	})
}
