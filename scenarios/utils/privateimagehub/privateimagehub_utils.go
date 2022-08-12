package privateimagehub

import (
	"fmt"

	applicationclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	"github.com/equinor/radix-cicd-canary/generated-client/models"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
)

// PasswordSet Checks if password is set
func PasswordSet(env envUtil.Env, appName string) error {
	return verifyStatus(env, appName, "Consistent")
}

// PasswordNotSet Verify that the private image hub password is not set
func PasswordNotSet(env envUtil.Env, appName string) error {
	return verifyStatus(env, appName, "Pending")
}

func verifyStatus(env envUtil.Env, appName, expectStatus string) error {
	imageHubs, err := List(env, appName)
	if err != nil {
		return err
	}
	imageHub := imageHubs[0]

	if imageHub.Status != expectStatus {
		return fmt.Errorf("private image hub status is %s, expected %s", imageHub.Status, expectStatus)
	}
	return nil
}

// SetPassword Sets password
func SetPassword(env envUtil.Env, appName string) error {
	imageHubs, err := List(env, appName)
	if err != nil {
		return err
	}
	imageHub := imageHubs[0]

	secretValue := env.GetPrivateImageHubPassword()
	secretParameters := models.SecretParameters{
		SecretValue: &secretValue,
	}

	params := applicationclient.NewUpdatePrivateImageHubsSecretValueParams().
		WithAppName(appName).
		WithServerName(*imageHub.Server).
		WithImageHubSecret(&secretParameters).
		WithImpersonateUser(env.GetImpersonateUserPointer()).
		WithImpersonateGroup(env.GetImpersonateGroupPointer())

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	_, err = client.UpdatePrivateImageHubsSecretValue(params, clientBearerToken)
	return err
}

// List Lists hubs
func List(env envUtil.Env, appName string) ([]*models.ImageHubSecret, error) {
	params := applicationclient.NewGetPrivateImageHubsParams().
		WithAppName(appName).
		WithImpersonateUser(env.GetImpersonateUserPointer()).
		WithImpersonateGroup(env.GetImpersonateGroupPointer())

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	privateImageHub, err := client.GetPrivateImageHubs(params, clientBearerToken)
	return privateImageHub.Payload, err
}
