package register

import (
	"context"
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
)

// ApplicationWithNoDeployKey Tests that we are able to register application
// with no deploy key and that deploy key is generated
func ApplicationWithNoDeployKey(ctx context.Context, cfg config.Config) error {
	appName := defaults.App1Name
	appRepo := defaults.App1Repository
	appCreator := defaults.App1Creator
	appConfigBranch := defaults.App1ConfigBranch
	appConfigurationItem := defaults.App1ConfigurationItem

	err := application.DeleteIfExist(ctx, cfg, appName)
	if err != nil {
		return err
	}

	_, err = application.Register(ctx, cfg, appName, appRepo, appCreator, appConfigBranch, appConfigurationItem, cfg.GetAppAdminGroup(), []string{cfg.GetAppReaderGroup()})
	if err != nil {
		return fmt.Errorf("failed to register application %s: %w", appName, err)
	}

	err = test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		return application.IsDefined(ctx, cfg, appName)
	})
	if err != nil {
		return err
	}

	if err := test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		return application.IsSharedSecretDefined(ctx, cfg, appName)
	}); err != nil {
		return err
	}

	return test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		return application.IsDeployKeyDefined(ctx, cfg, appName)
	})
}
