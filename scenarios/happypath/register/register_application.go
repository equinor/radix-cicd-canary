package register

import (
	"context"

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

	err = application.RegenerateDeployKey(ctx, cfg, appName, cfg.GetPrivateKey(), "")
	if err != nil {
		return errors.Wrapf(err, "failed to regenerate deploy key for application %s", appName)
	}

	return test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		return application.HasDeployKey(ctx, cfg, appName, cfg.GetPublicKey())
	})
}
