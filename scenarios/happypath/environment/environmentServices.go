package environment

import (
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/client/environment"
	"github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
)

// GetEnvironment gets an application by appName
func GetEnvironment(env env.Env, appName, envName string) (*models.Environment, error) {
	params := environmentclient.NewGetEnvironmentParams().
		WithAppName(appName).
		WithEnvName(appName).
		WithImpersonateUser(env.GetImpersonateUserPointer()).
		WithImpersonateGroup(env.GetImpersonateGroupPointer())
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetEnvironmentClient(env)

	result, err := client.GetEnvironment(params, clientBearerToken)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}
