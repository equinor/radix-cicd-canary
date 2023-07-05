package environment

import (
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/environment"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
)

// GetEnvironment gets an application by appName
func GetEnvironment(cfg config.Config, appName, envName string) (*models.Environment, error) {
	params := environmentclient.NewGetEnvironmentParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithImpersonateUser(cfg.GetImpersonateUser()).
		WithImpersonateGroup(cfg.GetImpersonateGroups())
	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetEnvironmentClient(cfg)

	result, err := client.GetEnvironment(params, clientBearerToken)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}
