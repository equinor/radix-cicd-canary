package register

import (
	"context"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/pkg/errors"
)

// ApplicationWithMainConfigBranch Tests that we are able to register the application
func ApplicationWithMainConfigBranch(ctx context.Context, cfg config.Config) error {
	appName := defaults.App4Name
	appRepo := defaults.App4Repository
	appSharedSecret := defaults.App4SharedSecret
	appCreator := defaults.App4Creator
	appConfigBranch := defaults.App4ConfigBranch
	appConfigurationItem := defaults.App4ConfigurationItem

	err := application.DeleteIfExist(ctx, cfg, appName)
	if err != nil {
		return err
	}

	_, err = application.Register(ctx, cfg, appName, appRepo, appSharedSecret, appCreator, appConfigBranch, appConfigurationItem, cfg.GetAppAdminGroup(), []string{cfg.GetAppReaderGroup()})
	if err != nil {
		return errors.Wrapf(err, "failed to register application %s", appName)
	}

	err = test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		return application.IsDefined(ctx, cfg, appName)
	})
	if err != nil {
		return err
	}

	err = application.RegenerateDeployKey(ctx, cfg, appName, cfg.GetPrivateKeyCanary4(), "")
	if err != nil {
		return err
	}

	return test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		return application.HasDeployKey(ctx, cfg, appName, cfg.GetPublicKeyCanary4())
	})
}
