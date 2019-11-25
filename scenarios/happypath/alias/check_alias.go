package alias

import (
	"fmt"
	"net/http"

	applicationclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/client/environment"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

const publicDomainNameEnvironmentVariable = "RADIX_PUBLIC_DOMAIN_NAME"

// DefaultResponding Checks if default alias of application is responding
func DefaultResponding(env envUtil.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	ok, _ := test.WaitForCheckFuncOrTimeout(env, isAppAliasDefined)
	publicDomainName := GetPublicDomainName(env, config.App2Component1Name)
	if publicDomainName == "" {
		return false, fmt.Errorf("Public domain name of alias is empty")
	}

	ok, _ = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) { return isAliasResponding(env, publicDomainName) })
	return ok, nil
}

func isAppAliasDefined(env envUtil.Env) (bool, interface{}) {
	appAlias := getApplicationAlias(env)
	if appAlias != nil {
		logger.Infof("App alias is defined %s. Now we can try to hit it to see if it responds", *appAlias)
		return true, *appAlias
	}

	logger.Info("App alias is not yet defined")
	return false, nil
}

func getApplicationAlias(env envUtil.Env) *string {
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

// GetPublicDomainName returns domain name for a component
func GetPublicDomainName(env envUtil.Env, forComponentName string) string {
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
			if componentName == forComponentName {
				return component.Variables[publicDomainNameEnvironmentVariable]
			}
		}
	}

	return ""
}

func isAliasResponding(env envUtil.Env, url string) (bool, interface{}) {
	req := httpUtils.CreateRequest(env, url, "GET", nil)
	client := http.DefaultClient
	resp, err := client.Do(req)

	if err == nil && resp.StatusCode == 200 {
		logger.Info("App alias responded ok")
		return true, nil
	}

	logger.Info("Alias is still not responding")
	return false, nil
}
