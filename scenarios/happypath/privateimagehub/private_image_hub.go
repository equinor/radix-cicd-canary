package privateimagehub

import (
	"context"
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/happypath/environment"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/privateimagehub"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/rs/zerolog/log"
)

// Set runs tests related to private image hub. Expect canary2 to be built and deployed before test run
func Set(ctx context.Context, cfg config.Config) error {
	appName := defaults.App2Name
	appCtx := log.Ctx(ctx).With().Str("app", appName).Logger().WithContext(ctx)

	err := privateimagehub.PasswordNotSet(cfg, appName)
	if err != nil {
		return err
	}
	log.Ctx(appCtx).Info().Msg("SUCCESS: private image hub is not set")

	err = test.WaitForCheckFuncOrTimeout(appCtx, cfg, func(cfg config.Config, ctx context.Context) error {
		return podNotLoaded(cfg, appName)
	})
	if err != nil {
		return fmt.Errorf("%s component is running before private image hub password was set. %v", defaults.App2ComponentPrivateImageHubName, err)
	}
	log.Ctx(appCtx).Info().Msg("SUCCESS: container is not loaded")

	err = privateimagehub.SetPassword(cfg, appName)
	if err != nil {
		return fmt.Errorf("failed to set private image hub password. %v", err)
	}
	log.Ctx(appCtx).Info().Msg("SUCCESS: set private image hub password")

	err = test.WaitForCheckFuncOrTimeout(appCtx, cfg, func(cfg config.Config, ctx context.Context) error {
		return podLoaded(cfg, appName)
	})
	if err != nil {
		return fmt.Errorf("%s component does not run after setting private image hub password. Error %v", defaults.App2ComponentPrivateImageHubName, err.Error())
	}
	log.Ctx(appCtx).Info().Msg("SUCCESS: container is loaded with updated image hub password")

	err = privateimagehub.PasswordSet(cfg, appName)
	if err != nil {
		return err
	}
	log.Ctx(appCtx).Info().Msg("SUCCESS: private image hub is verified set")

	return nil
}

func podNotLoaded(cfg config.Config, appName string) error {
	return verifyPrivateImageHubPodStatus(cfg, appName, "Failing")
}

func podLoaded(cfg config.Config, appName string) error {
	return verifyPrivateImageHubPodStatus(cfg, appName, "Running")
}

func verifyPrivateImageHubPodStatus(cfg config.Config, appName string, expectedStatus string) error {
	actualStatus, err := getPrivateImageHubComponentStatus(cfg, appName)
	if err != nil {
		return err
	}
	if actualStatus != expectedStatus {
		return fmt.Errorf("expected status %s on component %s - was %s", expectedStatus, defaults.App2ComponentPrivateImageHubName, actualStatus)
	}
	return nil
}

func getPrivateImageHubComponentStatus(cfg config.Config, appName string) (string, error) {
	envQA, err := environment.GetEnvironment(cfg, appName, defaults.App2EnvironmentName)
	if err != nil {
		return "", err
	}
	for _, comp := range envQA.ActiveDeployment.Components {
		if *comp.Name == defaults.App2ComponentPrivateImageHubName && len(comp.ReplicaList) > 0 {
			if replica := comp.ReplicaList[0]; replica != nil {
				return *replica.ReplicaStatus.Status, nil
			}
		}
	}
	return "", nil
}
