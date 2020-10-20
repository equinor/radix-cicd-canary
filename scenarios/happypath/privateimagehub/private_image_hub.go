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

var logger *log.Entry

// Set runs tests related to private image hub. Expect canary2 to be built and deployed before test run
func Set(env envUtil.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	err := privateimagehub.PasswordNotSet(env, config.App2Name)
	if err != nil {
		return false, err
	}
	logger.Infof("SUCCESS: private image hub is not set")

	ok, errorMessage := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		errorMessage := podNotLoaded(env)
		if len(errorMessage) > 0 {
			return false, err
		}
		return true, nil
	})
	if !ok {
		return false, fmt.Errorf("%s component is running before private image hub password was se. %s", config.App2ComponentPrivateImageHubName, errorMessage)
	}
	logger.Infof("SUCCESS: container is not loaded")

	err = privateimagehub.SetPassword(env, config.App2Name)
	if err != nil {
		return false, fmt.Errorf("Failed to set private image hub password. %v", err)
	}
	logger.Infof("SUCCESS: set private image hub password")

	ok, _ = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		errorMessage := podLoaded(env)
		if len(errorMessage) > 0 {
			return false, errorMessage
		}
		return true, nil
	})
	logger.Infof("SUCCESS: container is loaded")
	if !ok {
		return false, fmt.Errorf("%s component does not run after setting private image hub password. Error %s", config.App2ComponentPrivateImageHubName, errorMessage)
	}

	err = privateimagehub.PasswordSet(env, config.App2Name)
	if err != nil {
		return false, err
	}
	logger.Infof("SUCCESS: private image hub is verified set")

	return true, nil
}

func podNotLoaded(env envUtil.Env) string {
	return verifyPrivateImageHubPodStatus(env, "Failing")
}

func podLoaded(env envUtil.Env) string {
	return verifyPrivateImageHubPodStatus(env, "Running")
}

func verifyPrivateImageHubPodStatus(env envUtil.Env, expectedStatus string) string {
	actualStatus, err := getPrivateImageHubComponentStatus(env)
	if err != nil {
		return err.Error()
	} else if actualStatus != expectedStatus {
		logger.Debugf("expected status %s on component %s - was %s", expectedStatus, config.App2ComponentPrivateImageHubName, actualStatus)
		return fmt.Sprintf("expected status %s on component %s - was %s", expectedStatus, config.App2ComponentPrivateImageHubName, actualStatus)
	}
	return ""
}

func getPrivateImageHubComponentStatus(env envUtil.Env) (string, error) {
	appName := config.App2Name
	envQA, err := environment.GetEnvironment(env, appName, config.App2EnvironmentName)
	if err != nil {
		return "", err
	}
	for _, comp := range envQA.ActiveDeployment.Components {
		if *comp.Name == config.App2ComponentPrivateImageHubName && comp.ReplicaList != nil {
			replica := comp.ReplicaList[0]
			return *replica.ReplicaStatus.Status, nil
		}
	}
	return "", nil
}
