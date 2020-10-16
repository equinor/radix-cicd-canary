package application

import (
	"net/http"
	"strings"

	applicationclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/client/environment"
	apiclient "github.com/equinor/radix-cicd-canary/generated-client/client/platform"
	"github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

const canonicalDomainNameEnvironmentVariable = "RADIX_CANONICAL_DOMAIN_NAME"
const publicDomainNameEnvironmentVariable = "RADIX_PUBLIC_DOMAIN_NAME"

// Register Will register application
func Register(env env.Env, appName, appRepo, appSharedSecret, appOwner, appCreator, publicKey, privateKey string, wbs string) (*apiclient.RegisterApplicationOK, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()
	bodyParameters := models.ApplicationRegistration{
		Name:         &appName,
		Repository:   &appRepo,
		SharedSecret: &appSharedSecret,
		Owner:        &appOwner,
		Creator:      &appCreator,
		AdGroups:     nil,
		PublicKey:    publicKey,
		PrivateKey:   privateKey,
		WBS:          wbs,
	}

	params := apiclient.NewRegisterApplicationParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithApplicationRegistration(&bodyParameters)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetPlatformClient(env)

	return client.RegisterApplication(params, clientBearerToken)
}

// Delete Deletes application
func Delete(env env.Env, appName string) (bool, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := applicationclient.NewDeleteApplicationParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(appName)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	_, err := client.DeleteApplication(params, clientBearerToken)
	if err != nil {
		log.Errorf("Error calling DeleteApplication for application %s: %v", appName, err)
	}

	return err == nil, err
}

// Deploy Deploys application
func Deploy(env env.Env, appName, toEnvironment string) (*applicationclient.TriggerPipelineDeployOK, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	bodyParameters := models.PipelineParametersDeploy{
		ToEnvironment: toEnvironment,
	}

	params := applicationclient.NewTriggerPipelineDeployParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(appName).
		WithPipelineParametersDeploy(&bodyParameters)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	return client.TriggerPipelineDeploy(params, clientBearerToken)
}

// IsDefined Checks if application is defined
func IsDefined(env env.Env, appName string) (bool, interface{}) {
	_, err := Get(env, appName)
	if err == nil {
		return true, nil
	}

	log.Infof("Application %s is not defined", appName)
	return false, nil
}

// Get gets an application by appName
func Get(env env.Env, appName string) (*models.Application, error) {
	params := applicationclient.NewGetApplicationParams().
		WithAppName(appName).
		WithImpersonateUser(env.GetImpersonateUserPointer()).
		WithImpersonateGroup(env.GetImpersonateGroupPointer())
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	result, err := client.GetApplication(params, clientBearerToken)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}

// IsAliasDefined Checks if app alias is defined
func IsAliasDefined(env env.Env, appName string) (bool, interface{}) {
	appAlias := getAlias(env, appName)
	if appAlias != nil {
		log.Infof("App alias is defined %s. Now we can try to hit it to see if it responds", *appAlias)
		return true, *appAlias
	}

	log.Info("App alias is not yet defined")
	return false, nil
}

func getAlias(env env.Env, appName string) *string {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := applicationclient.NewGetApplicationParams().
		WithAppName(appName).
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

// IsRunningInActiveCluster Check if app is running in active cluster
func IsRunningInActiveCluster(publicDomainName, canonicalDomainName string) bool {
	if strings.EqualFold(publicDomainName, canonicalDomainName) {
		return false
	}

	return true
}

// TryGetPublicDomainName Waits for public domain name to be defined
func TryGetPublicDomainName(env env.Env, appName, environmentName, componentName string) (bool, interface{}) {
	publicDomainName := getEnvVariable(env, appName, environmentName, componentName, publicDomainNameEnvironmentVariable)
	if publicDomainName == "" {
		return false, nil
	}
	return true, publicDomainName
}

// TryGetCanonicalDomainName Waits for canonical domain name to be defined
func TryGetCanonicalDomainName(env env.Env, appName, environmentName, componentName string) (bool, interface{}) {
	canonicalDomainName := getEnvVariable(env, appName, environmentName, componentName, publicDomainNameEnvironmentVariable)
	if canonicalDomainName == "" {
		return false, nil
	}

	return true, canonicalDomainName
}

func getEnvVariable(env env.Env, appName, envName, forComponentName, variableName string) string {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := environmentclient.NewGetEnvironmentParams().
		WithAppName(appName).
		WithEnvName(envName).
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
				return component.Variables[variableName]
			}
		}
	}

	return ""
}

// AreResponding Checks if all endpoint responds
func AreResponding(env env.Env, urls ...string) (bool, interface{}) {
	for _, url := range urls {
		ok, _ := IsResponding(env, url)
		if !ok {
			return false, nil
		}
	}

	return true, nil
}

// IsResponding Checks if endpoint is responding
func IsResponding(env env.Env, url string) (bool, interface{}) {
	req := httpUtils.CreateRequest(env, url, "GET", nil)
	client := http.DefaultClient
	resp, err := client.Do(req)

	if err == nil && resp.StatusCode == 200 {
		log.Info("App alias responded ok")
		return true, nil
	}

	if err != nil {
		log.Debugf("Request to alias '%s' returned error %v", url, err)
	}

	if resp != nil {
		log.Debugf("Request to alias '%s' returned status %v", url, resp.StatusCode)
	}

	if err == nil && resp == nil {
		log.Debugf("Request to alias '%s' returned no response and no err.", url)
	}

	log.Infof("Alias '%s' is still not responding", url)
	return false, nil
}
