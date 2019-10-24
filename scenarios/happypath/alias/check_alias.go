package alias

import (
	"net/http"

	applicationclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/client/environment"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

const publicDomainNameEnvironmentVariable = "RADIX_PUBLIC_DOMAIN_NAME"

// DefaultResponding Checks if default alias of application is responding
func DefaultResponding(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	ok, _ := test.WaitForCheckFunc(env, isAppAliasDefined)
	publicDomainName := getPublicDomainName(env)

	ok, _ = test.WaitForCheckFuncWithArguments(env, isAliasResponding, []string{publicDomainName})
	return ok, nil
}

func isAppAliasDefined(env env.Env, args []string) (bool, interface{}) {
	appAlias := getApplicationAlias(env)
	if appAlias != nil {
		logger.Info("App alias is defined. Now we can try to hit it to see if it responds")
		return true, *appAlias
	}

	logger.Info("App alias is not yet defined")
	return false, nil
}

func getApplicationAlias(env env.Env) *string {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := applicationclient.NewGetApplicationParams().
		WithAppName(config.App2Name).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	applicationDetails, err := client.GetApplication(params, clientBearerToken)
	if err == nil && applicationDetails.Payload != nil && applicationDetails.Payload.AppAlias != nil {
		return applicationDetails.Payload.AppAlias.URL
	}

	return nil
}

func getPublicDomainName(env env.Env) string {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := environmentclient.NewGetEnvironmentParams().
		WithAppName(config.App2Name).
		WithEnvName(config.App2EnvironmentName).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetEnvironmentClient(env)

	environmentDetails, err := client.GetEnvironment(params, clientBearerToken)
	if err == nil &&
		environmentDetails.Payload != nil &&
		environmentDetails.Payload.ActiveDeployment != nil {
		for _, component := range environmentDetails.Payload.ActiveDeployment.Components {
			componentName := *component.Name
			if componentName == config.App2Component1Name {
				return component.Variables[publicDomainNameEnvironmentVariable]
			}
		}
	}

	return ""
}

func isAliasResponding(env env.Env, args []string) (bool, interface{}) {
	req := httpUtils.CreateRequest(env, args[0], "GET", nil)
	client := http.DefaultClient
	resp, err := client.Do(req)

	if err == nil && resp.StatusCode == 200 {
		logger.Info("App alias responded ok")
		return true, nil
	}

	logger.Info("Alias is still not responding")
	return false, nil
}
