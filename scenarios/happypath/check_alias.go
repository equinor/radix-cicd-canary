package happypath

import (
	"net/http"

	applicationclient "github.com/equinor/radix-cicd-canary-golang/generated-client/client/application"
	environmentclient "github.com/equinor/radix-cicd-canary-golang/generated-client/client/environment"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary-golang/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

const (
	aliasRespondingTestName             = "DefaultAliasResponding"
	publicDomainNameEnvironmentVariable = "RADIX_PUBLIC_DOMAIN_NAME"
)

func defaultAliasResponding() string {
	ok, _ := test.WaitForCheckFunc(isAppAliasDefined)
	publicDomainName := getPublicDomainName()

	ok, _ = test.WaitForCheckFuncWithArguments(isAliasResponding, []string{publicDomainName})

	if ok {
		addTestSuccess(aliasRespondingTestName)
	} else {
		addTestError(aliasRespondingTestName)
	}

	return aliasRespondingTestName
}

func isAppAliasDefined(args []string) (bool, interface{}) {
	appAlias := getApplicationAlias()
	if appAlias != nil {
		log.Info("App alias is defined. Now we can try to hit it to see if it responds")
		return true, *appAlias
	}

	log.Info("App alias is not yet defined")
	return false, nil
}

func getApplicationAlias() *string {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := applicationclient.NewGetApplicationParams().
		WithAppName(app2Name).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken()
	client := httpUtils.GetApplicationClient()

	applicationDetails, err := client.GetApplication(params, clientBearerToken)
	if err == nil {
		return applicationDetails.Payload.AppAlias.URL
	}

	return nil
}

func getPublicDomainName() string {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := environmentclient.NewGetEnvironmentParams().
		WithAppName(app2Name).
		WithEnvName(app2EnvironmentName).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken()
	client := httpUtils.GetEnvironmentClient()

	environmentDetails, err := client.GetEnvironment(params, clientBearerToken)
	if err == nil && environmentDetails.Payload != nil {
		for _, component := range environmentDetails.Payload.ActiveDeployment.Components {
			componentName := *component.Name
			if componentName == app2Component1Name {
				return component.Variables[publicDomainNameEnvironmentVariable]
			}
		}
	}

	return ""
}

func isAliasResponding(args []string) (bool, interface{}) {
	req := httpUtils.CreateRequest(args[0], "GET", nil)
	client := http.DefaultClient
	resp, err := client.Do(req)

	if err == nil && resp.StatusCode == 200 {
		log.Info("App alias responded ok")
		return true, nil
	}

	log.Info("Alias is still not responding")
	return false, nil
}
