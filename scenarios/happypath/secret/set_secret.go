package secret

import (
	"context"
	"fmt"

	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/environment"
	models "github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger zerolog.Logger

// Set Test that we are able to set secret
func Set(ctx context.Context, cfg config.Config, suiteName string) error {

	err := test.WaitForCheckFuncOrTimeout(cfg, isDeploymentConsistent, ctx)
	if err != nil {
		return err
	}

	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := environmentclient.NewChangeComponentSecretParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithAppName(defaults.App2Name).
		WithEnvName(defaults.App2EnvironmentName).
		WithComponentName(defaults.App2Component2Name).
		WithSecretName(defaults.App2SecretName).
		WithComponentSecret(
			&models.SecretParameters{
				SecretValue: stringPtr(defaults.App2SecretValue),
			})

	client := httpUtils.GetEnvironmentClient(cfg)
	_, err = client.ChangeComponentSecret(params, nil)
	if err != nil {
		return fmt.Errorf("error calling ChangeComponentSecret for application %s: %v", defaults.App2Name, err)
	}
	return nil
}

func isDeploymentConsistent(cfg config.Config, ctx context.Context) error {
	environmentDetails := getEnvironmentDetails(cfg)
	if environmentDetails != nil &&
		environmentDetails.ActiveDeployment != nil &&
		environmentDetails.Status != "" &&
		len(environmentDetails.Secrets) > 0 {
		log.Ctx(ctx).Info().Msg("Deployment is consistent. We can set the secret.")
		return nil
	}
	return fmt.Errorf("deployment is not consistent")
}

func getEnvironmentDetails(cfg config.Config) *models.Environment {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := environmentclient.NewGetEnvironmentParams().
		WithAppName(defaults.App2Name).
		WithEnvName(defaults.App2EnvironmentName).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)
	client := httpUtils.GetEnvironmentClient(cfg)
	environmentDetails, err := client.GetEnvironment(params, nil)
	if err == nil && environmentDetails.Payload != nil {
		return environmentDetails.Payload
	}

	return nil
}

func stringPtr(str string) *string {
	return &str
}
