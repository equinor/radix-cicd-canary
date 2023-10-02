package register

import (
	"context"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/rs/zerolog/log"
)

// Application Tests that we are able to register application
// with deploy key set
func Application(ctx context.Context, cfg config.Config, suiteName string) error {
	appName := defaults.App3Name
	appRepo := defaults.App3Repository
	appSharedSecret := defaults.App3SharedSecret
	appCreator := defaults.App3Creator
	appConfigurationItem := defaults.App3ConfigurationItem
	appConfigBranch := defaults.App3ConfigBranch
	appCtx := log.Ctx(ctx).With().Str("app", appName).Logger().WithContext(ctx)

	err := application.DeleteIfExist(cfg, appName, appCtx)
	if err != nil {
		return err
	}

	_, err = application.Register(cfg, appName, appRepo, appSharedSecret, appCreator, appConfigBranch, appConfigurationItem, cfg.GetAppAdminGroup(), []string{cfg.GetAppReaderGroup()})
	if err != nil {
		return err
	}

	err = test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config, ctx context.Context) error {
		return application.IsDefined(ctx, cfg, defaults.App3Name)
	}, appCtx)
	if err != nil {
		return err
	}

	err = application.RegenerateDeployKey(cfg, appName, cfg.GetPrivateKeyCanary3(), "", appCtx)
	if err != nil {
		return err
	}

	return test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config, ctx context.Context) error {
		return application.HasDeployKey(cfg, appName, cfg.GetPublicKeyCanary3(), ctx)
	}, appCtx)
}
