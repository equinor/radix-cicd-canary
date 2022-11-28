package secret

import (
	"fmt"

	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/client/environment"
	models "github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Set Test that we are able to set secret
func Set(cfg config.Config, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	err := test.WaitForCheckFuncOrTimeout(cfg, isDeploymentConsistent, logger)
	if err != nil {
		return err
	}

	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroup()

	params := environmentclient.NewChangeComponentSecretParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(defaults.App2Name).
		WithEnvName(defaults.App2EnvironmentName).
		WithComponentName(defaults.App2Component2Name).
		WithSecretName(defaults.App2SecretName).
		WithComponentSecret(
			&models.SecretParameters{
				SecretValue: stringPtr(defaults.App2SecretValue),
			})

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetEnvironmentClient(cfg)

	_, err = client.ChangeComponentSecret(params, clientBearerToken)
	if err != nil {
		return fmt.Errorf("error calling ChangeComponentSecret for application %s: %v", defaults.App2Name, err)
	}
	return nil
}

func isDeploymentConsistent(cfg config.Config) error {
	environmentDetails := getEnvironmentDetails(cfg)
	if environmentDetails != nil &&
		environmentDetails.ActiveDeployment != nil &&
		environmentDetails.Status != "" &&
		len(environmentDetails.Secrets) > 0 {
		logger.Info("Deployment is consistent. We can set the secret.")
		return nil
	}
	return fmt.Errorf("deployment is not consistent")
}

func getEnvironmentDetails(cfg config.Config) *models.Environment {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroup()

	params := environmentclient.NewGetEnvironmentParams().
		WithAppName(defaults.App2Name).
		WithEnvName(defaults.App2EnvironmentName).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetEnvironmentClient(cfg)

	environmentDetails, err := client.GetEnvironment(params, clientBearerToken)
	if err == nil && environmentDetails.Payload != nil {
		return environmentDetails.Payload
	}

	return nil
}

func stringPtr(str string) *string {
	return &str
}
