package privateimagehub

import (
	"context"
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/component"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/privateimagehub"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/rs/zerolog/log"
)

// Set runs tests related to private image hub. Expect canary2 to be built and deployed before test run
func Set(ctx context.Context, cfg config.Config) error {
	if err := privateimagehub.PasswordNotSet(cfg, defaults.App3Name); err != nil {
		return err
	}
	log.Ctx(ctx).Info().Msg("verified private image hub is not set")

	log.Ctx(ctx).Info().Msg("verify that all replicas are in failing state")
	err := test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		return allReplicasFailing(ctx, cfg)
	})
	if err != nil {
		return fmt.Errorf("failed to verify that all replicas are in failing state: %w", err)
	}
	log.Ctx(ctx).Info().Msg("verified that all replicas are in failing state")

	log.Ctx(ctx).Debug().Str("app", defaults.App3Name).Msg("set privateimagehub password for application")
	err = privateimagehub.SetPassword(cfg, defaults.App3Name)
	if err != nil {
		return fmt.Errorf("failed to set private image hub password: %w", err)
	}
	log.Ctx(ctx).Info().Msg("successfully set private image hub password")

	err = privateimagehub.PasswordSet(cfg, defaults.App3Name)
	if err != nil {
		return err
	}
	log.Ctx(ctx).Info().Msg("verified private image hub password is set")

	log.Ctx(ctx).Info().Msg("verify that all replicas are in running state")
	err = test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		return allReplicasRunning(ctx, cfg)
	})
	if err != nil {
		return fmt.Errorf("failed to verify that all replicas are in running state: %w", err)
	}
	log.Ctx(ctx).Info().Msg("verified all replicas are in running state")

	return nil
}

func allReplicasFailing(ctx context.Context, cfg config.Config) error {
	return verifyPrivateImageHubPodStatus(ctx, cfg, "Failing")
}

func allReplicasRunning(ctx context.Context, cfg config.Config) error {
	return verifyPrivateImageHubPodStatus(ctx, cfg, "Running")
}

func verifyPrivateImageHubPodStatus(ctx context.Context, cfg config.Config, expectedStatus string) error {
	return component.AllReplicasHaveExpectedStatus(ctx, cfg, defaults.App3Name, defaults.App3EnvironmentName, defaults.App3Component1Name, expectedStatus)
}
