package privateimagehub

import (
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/happypath/environment"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/privateimagehub"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

// Set runs tests related to private image hub. Expect canary2 to be built and deployed before test run
func Set(env envUtil.Env, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})

	err := privateimagehub.PasswordNotSet(env, config.App2Name)
	if err != nil {
		return err
	}
	logger.Infof("SUCCESS: private image hub is not set")

	err = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) error {
		return podNotLoaded(env)
	}, logger)
	if err != nil {
		return fmt.Errorf("%s component is running before private image hub password was set. %v", config.App2ComponentPrivateImageHubName, err)
	}
	logger.Infof("SUCCESS: container is not loaded")

	err = privateimagehub.SetPassword(env, config.App2Name)
	if err != nil {
		return fmt.Errorf("failed to set private image hub password. %v", err)
	}
	logger.Infof("SUCCESS: set private image hub password")

	err = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) error {
		return podLoaded(env)
	}, logger)
	if err != nil {
		return fmt.Errorf("%s component does not run after setting private image hub password. Error %v", config.App2ComponentPrivateImageHubName, err.Error())
	}
	logger.Infof("SUCCESS: container is loaded with updated image hub password")

	err = privateimagehub.PasswordSet(env, config.App2Name)
	if err != nil {
		return err
	}
	logger.Infof("SUCCESS: private image hub is verified set")

	return nil
}

func podNotLoaded(env envUtil.Env) error {
	return verifyPrivateImageHubPodStatus(env, "Failing")
}

func podLoaded(env envUtil.Env) error {
	return verifyPrivateImageHubPodStatus(env, "Running")
}

func verifyPrivateImageHubPodStatus(env envUtil.Env, expectedStatus string) error {
	actualStatus, err := getPrivateImageHubComponentStatus(env)
	if err != nil {
		return err
	}
	if actualStatus != expectedStatus {
		return fmt.Errorf("expected status %s on component %s - was %s", expectedStatus, config.App2ComponentPrivateImageHubName, actualStatus)
	}
	return nil
}

func getPrivateImageHubComponentStatus(env envUtil.Env) (string, error) {
	appName := config.App2Name
	envQA, err := environment.GetEnvironment(env, appName, config.App2EnvironmentName)
	if err != nil {
		return "", err
	}
	for _, comp := range envQA.ActiveDeployment.Components {
		if *comp.Name == config.App2ComponentPrivateImageHubName && len(comp.ReplicaList) > 0 {
			if replica := comp.ReplicaList[0]; replica != nil {
				return *replica.ReplicaStatus.Status, nil
			}
		}
	}
	return "", nil
}
