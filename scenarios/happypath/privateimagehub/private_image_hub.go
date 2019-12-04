package privateimagehub

import (
	"fmt"

	applicationclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	"github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/environment"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// PrivateImageHub runs tests related to private image hub. Expect canary2 to be built and deployed before test run
func PrivateImageHub(env envUtil.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	err := privateImageHubPasswordNotSet(env)
	if err != nil {
		return false, err
	}
	logger.Infof("SUCCESS: private image hub is not set")

	ok, _ := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		err := podNotLoaded(env)
		if err != nil {
			return false, err
		}
		return true, nil
	})
	if !ok {
		return false, fmt.Errorf("%s component is running before private image hub password was se. %v", config.App2ComponentPrivateImageHubName, err)
	}
	logger.Infof("SUCCESS: container is not loaded")

	err = setPrivateImageHubPassword(env)
	if err != nil {
		return false, fmt.Errorf("Failed to set private image hub password. %v", err)
	}
	logger.Infof("SUCCESS: set private image hub password")

	ok, _ = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		err := podLoaded(env)
		if err != nil {
			return false, err
		}
		return true, nil
	})
	logger.Infof("SUCCESS: container is loaded")
	if !ok {
		return false, fmt.Errorf("%s component does not run after setting private image hub password. Error %v", config.App2ComponentPrivateImageHubName, err)
	}

	err = privateImageHubPasswordSet(env)
	if err != nil {
		return false, err
	}
	logger.Infof("SUCCESS: private image hub is verified set")

	return true, nil
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
	} else if actualStatus != expectedStatus {
		logger.Errorf("expected status %s on component %s - was %s", expectedStatus, config.App2ComponentPrivateImageHubName, actualStatus)
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
		if *comp.Name == config.App2ComponentPrivateImageHubName && comp.ReplicaList != nil {
			replica := comp.ReplicaList[0]
			return *replica.ReplicaStatus.Status, nil
		}
	}
	return "", nil
}

func privateImageHubPasswordSet(env envUtil.Env) error {
	return verifyPrivateImageHubStatus(env, "Consistent")
}

func privateImageHubPasswordNotSet(env envUtil.Env) error {
	return verifyPrivateImageHubStatus(env, "Pending")
}

func verifyPrivateImageHubStatus(env envUtil.Env, expectStatus string) error {
	imageHubs, err := getPrivateImageHubs(env, config.App2Name)
	if err != nil {
		return err
	}
	imageHub := imageHubs[0]

	if imageHub.Status != expectStatus {
		return fmt.Errorf("Private image hub status is %s, expected %s", imageHub.Status, expectStatus)
	}
	return nil
}

func setPrivateImageHubPassword(env envUtil.Env) error {
	imageHubs, err := getPrivateImageHubs(env, config.App2Name)
	if err != nil {
		return err
	}
	imageHub := imageHubs[0]

	secretValue := env.GetPrivateImageHubPassword()
	secretParameters := models.SecretParameters{
		SecretValue: &secretValue,
	}

	params := applicationclient.NewUpdatePrivateImageHubsSecretValueParams().
		WithAppName(config.App2Name).
		WithServerName(*imageHub.Server).
		WithImageHubSecret(&secretParameters).
		WithImpersonateUser(env.GetImpersonateUserPointer()).
		WithImpersonateGroup(env.GetImpersonateGroupPointer())

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	_, err = client.UpdatePrivateImageHubsSecretValue(params, clientBearerToken)
	return err
}

func getPrivateImageHubs(env envUtil.Env, appName string) ([]*models.ImageHubSecret, error) {
	params := applicationclient.NewGetPrivateImageHubsParams().
		WithAppName(appName).
		WithImpersonateUser(env.GetImpersonateUserPointer()).
		WithImpersonateGroup(env.GetImpersonateGroupPointer())

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	privateImageHub, err := client.GetPrivateImageHubs(params, clientBearerToken)
	return privateImageHub.Payload, err
}
