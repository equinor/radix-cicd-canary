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

// ApplicationWithNoDeployKey Tests that we are able to register application
// with no deploy key and that deploy key is generated
func ApplicationWithNoDeployKey(ctx context.Context, cfg config.Config) error {
	appName := defaults.App1Name
	appRepo := defaults.App1Repository
	appSharedSecret := defaults.App1SharedSecret
	appCreator := defaults.App1Creator
	appConfigBranch := defaults.App1ConfigBranch
	appConfigurationItem := defaults.App1ConfigurationItem
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

	return test.WaitForCheckFuncOrTimeout(appCtx, cfg, func(cfg config.Config, ctx context.Context) error {
		return application.IsDeployKeyDefined(ctx, cfg, appName)
	})
}
