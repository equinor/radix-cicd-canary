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

// Application Tests that we are able to register application
// with deploy key set
func Application(ctx context.Context, cfg config.Config, suiteName string) error {
	appName := defaults.App2Name
	appRepo := defaults.App2Repository
	appSharedSecret := defaults.App2SharedSecret
	appCreator := defaults.App2Creator
	appConfigBranch := defaults.App2ConfigBranch
	appConfigurationItem := defaults.App2ConfigurationItem
	appCtx := log.With().Str("suite", suiteName).Str("app", appName).Logger().WithContext(ctx)

	err := application.DeleteIfExist(cfg, appName, appCtx)
	if err != nil {
		return err
	}

	_, err = application.Register(cfg, appName, appRepo, appSharedSecret, appCreator, appConfigBranch, appConfigurationItem, cfg.GetAppAdminGroup(), []string{cfg.GetAppReaderGroup()})
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("failed to register application %s", appName))
	}

	err = test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config, ctx context.Context) error {
		return application.IsDefined(ctx, cfg, appName)
	}, appCtx)
	if err != nil {
		return err
	}

	err = application.RegenerateDeployKey(cfg, appName, cfg.GetPrivateKey(), "", appCtx)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("failed to regenerate deploy key for application %s", appName))
	}

	return test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config, ctx context.Context) error {
		return application.HasDeployKey(cfg, appName, cfg.GetPublicKey(), ctx)
	}, appCtx)
}
