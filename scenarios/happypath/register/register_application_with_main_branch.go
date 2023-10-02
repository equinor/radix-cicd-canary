package register

import (
	"context"
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// ApplicationWithMainConfigBranch Tests that we are able to register the application
func ApplicationWithMainConfigBranch(ctx context.Context, cfg config.Config, suiteName string) error {
	appName := defaults.App4Name
	appRepo := defaults.App4Repository
	appSharedSecret := defaults.App4SharedSecret
	appCreator := defaults.App4Creator
	appConfigBranch := defaults.App4ConfigBranch
	appConfigurationItem := defaults.App4ConfigurationItem
	appCtx := log.Ctx(ctx).With().Str("app", appName).Logger().WithContext(ctx)

	err := application.DeleteIfExist(appCtx, cfg, appName)
	if err != nil {
		return err
	}

	_, err = application.Register(appCtx, cfg, appName, appRepo, appSharedSecret, appCreator, appConfigBranch, appConfigurationItem, cfg.GetAppAdminGroup(), []string{cfg.GetAppReaderGroup()})
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("failed to register application %s", appName))
	}

	err = test.WaitForCheckFuncOrTimeout(appCtx, cfg, func(cfg config.Config, ctx context.Context) error {
		return application.IsDefined(ctx, cfg, appName)
	})
	if err != nil {
		return err
	}

	err = application.RegenerateDeployKey(appCtx, cfg, appName, cfg.GetPrivateKeyCanary4(), "")
	if err != nil {
		return err
	}

	return test.WaitForCheckFuncOrTimeout(appCtx, cfg, func(cfg config.Config, ctx context.Context) error {
		return application.HasDeployKey(ctx, cfg, appName, cfg.GetPublicKeyCanary4())
	})
}
