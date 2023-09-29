package privateimagehub

import (
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/happypath/environment"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/privateimagehub"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/rs/zerolog/log"
)

// Set runs tests related to private image hub. Expect canary2 to be built and deployed before test run
func Set(cfg config.Config, suiteName string) error {
	logger := log.With().Str("suite", suiteName).Logger()

	err := privateimagehub.PasswordNotSet(cfg, defaults.App2Name)
	if err != nil {
		return err
	}
	logger.Info().Msg("SUCCESS: private image hub is not set")

	err = test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return podNotLoaded(cfg)
	}, logger)
	if err != nil {
		return fmt.Errorf("%s component is running before private image hub password was set. %v", defaults.App2ComponentPrivateImageHubName, err)
	}
	logger.Info().Msg("SUCCESS: container is not loaded")

	err = privateimagehub.SetPassword(cfg, defaults.App2Name)
	if err != nil {
		return fmt.Errorf("failed to set private image hub password. %v", err)
	}
	logger.Info().Msg("SUCCESS: set private image hub password")

	err = test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return podLoaded(cfg)
	}, logger)
	if err != nil {
		return fmt.Errorf("%s component does not run after setting private image hub password. Error %v", defaults.App2ComponentPrivateImageHubName, err.Error())
	}
	logger.Info().Msg("SUCCESS: container is loaded with updated image hub password")

	err = privateimagehub.PasswordSet(cfg, defaults.App2Name)
	if err != nil {
		return err
	}
	logger.Info().Msg("SUCCESS: private image hub is verified set")

	return nil
}

func podNotLoaded(cfg config.Config) error {
	return verifyPrivateImageHubPodStatus(cfg, "Failing")
}

func podLoaded(cfg config.Config) error {
	return verifyPrivateImageHubPodStatus(cfg, "Running")
}

func verifyPrivateImageHubPodStatus(cfg config.Config, expectedStatus string) error {
	actualStatus, err := getPrivateImageHubComponentStatus(cfg)
	if err != nil {
		return err
	}
	if actualStatus != expectedStatus {
		return fmt.Errorf("expected status %s on component %s - was %s", expectedStatus, defaults.App2ComponentPrivateImageHubName, actualStatus)
	}
	return nil
}

func getPrivateImageHubComponentStatus(cfg config.Config) (string, error) {
	appName := defaults.App2Name
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
