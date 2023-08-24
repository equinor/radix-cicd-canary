package privateimagehub

import (
	"fmt"

	applicationclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
)

// PasswordSet Checks if password is set
func PasswordSet(cfg config.Config, appName string) error {
	return verifyStatus(cfg, appName, "Consistent")
}

// PasswordNotSet Verify that the private image hub password is not set
func PasswordNotSet(cfg config.Config, appName string) error {
	return verifyStatus(cfg, appName, "Pending")
}

func verifyStatus(cfg config.Config, appName, expectStatus string) error {
	imageHubs, err := List(cfg, appName)
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
func SetPassword(cfg config.Config, appName string) error {
	imageHubs, err := List(cfg, appName)
	if err != nil {
		return err
	}
	imageHub := imageHubs[0]

	secretValue := cfg.GetPrivateImageHubPassword()
	secretParameters := models.SecretParameters{
		SecretValue: &secretValue,
	}

	params := applicationclient.NewUpdatePrivateImageHubsSecretValueParams().
		WithAppName(appName).
		WithServerName(*imageHub.Server).
		WithImageHubSecret(&secretParameters).
		WithImpersonateUser(cfg.GetImpersonateUser()).
		WithImpersonateGroup(cfg.GetImpersonateGroups())

	client := httpUtils.GetApplicationClient(cfg)
	_, err = client.UpdatePrivateImageHubsSecretValue(params, nil)
	return err
}

// List Lists hubs
func List(cfg config.Config, appName string) ([]*models.ImageHubSecret, error) {
	params := applicationclient.NewGetPrivateImageHubsParams().
		WithAppName(appName).
		WithImpersonateUser(cfg.GetImpersonateUser()).
		WithImpersonateGroup(cfg.GetImpersonateGroups())

	client := httpUtils.GetApplicationClient(cfg)
	privateImageHub, err := client.GetPrivateImageHubs(params, nil)
	return privateImageHub.Payload, err
}
