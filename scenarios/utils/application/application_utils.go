package application

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	applicationclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/client/environment"
	apiclient "github.com/equinor/radix-cicd-canary/generated-client/client/platform"
	"github.com/equinor/radix-cicd-canary/generated-client/models"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	kubeUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/kubernetes"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

const (
	publicDomainNameEnvironmentVariable  = "RADIX_PUBLIC_DOMAIN_NAME"
	canonicalEndpointEnvironmentVariable = "RADIX_CANONICAL_DOMAIN_NAME"
)

// Register Will register application
func Register(env envUtil.Env, appName, appRepo, appSharedSecret, appCreator, publicKey, privateKey, configBranch, configurationItem string) (*apiclient.RegisterApplicationOK, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()
	bodyParameters := models.ApplicationRegistrationRequest{
		ApplicationRegistration: &models.ApplicationRegistration{
			Name:              &appName,
			Repository:        &appRepo,
			SharedSecret:      &appSharedSecret,
			Creator:           &appCreator,
			AdGroups:          nil,
			PublicKey:         publicKey,
			PrivateKey:        privateKey,
			ConfigBranch:      &configBranch,
			ConfigurationItem: configurationItem,
		},
	}

	params := apiclient.NewRegisterApplicationParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithApplicationRegistration(&bodyParameters)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetPlatformClient(env)

	return client.RegisterApplication(params, clientBearerToken)
}

// DeleteByImpersonatedUser Deletes an application by the impersonated user
func DeleteByImpersonatedUser(env envUtil.Env, appName string, logger *log.Entry) error {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()
	logger.Debugf("delete an application %s by the impersonamed user %s, group %s", appName, impersonateUser, impersonateGroup)

	params := applicationclient.NewDeleteApplicationParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(appName)

	return delete(env, appName, params)
}

// DeleteByServiceAccount an application by the service account
func DeleteByServiceAccount(env envUtil.Env, appName string, logger *log.Entry) error {
	logger.Debugf("delete an application %s by the service account", appName)

	params := applicationclient.NewDeleteApplicationParams().
		WithAppName(appName)

	return delete(env, appName, params)
}

func delete(env envUtil.Env, appName string, params *applicationclient.DeleteApplicationParams) error {
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	_, err := client.DeleteApplication(params, clientBearerToken)
	if err != nil {
		return fmt.Errorf("failed deleting the application %s: %v", appName, err)
	}
	return nil
}

// Deploy Deploy application
func Deploy(env envUtil.Env, appName, toEnvironment string) (*applicationclient.TriggerPipelineDeployOK, error) {
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
func IsDefined(env envUtil.Env, appName string) error {
	_, err := Get(env, appName)
	if err == nil {
		return nil
	}
	return fmt.Errorf("application %s is not defined", appName)
}

func appNamespacesDoNotExist(env envUtil.Env, appName string) error {
	nsList, err := kubeUtils.GetKubernetesClient().CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{
		LabelSelector: labels.Set{"radix-app": appName}.String(),
	})
	if err != nil {
		return err
	}
	if len(nsList.Items) > 0 {
		return fmt.Errorf("there are %d namespaces for the application %s", len(nsList.Items), appName)
	}
	return nil
}

// DeleteIfExist Delete application if it exists
func DeleteIfExist(env envUtil.Env, appName string, logger *log.Entry) error {
	err := IsDefined(env, appName)
	if err != nil {
		return nil
	}
	DeleteByServiceAccount(env, appName, logger)
	return test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) error {
		return appNamespacesDoNotExist(env, appName)
	}, logger)
}

// Get gets an application by appName
func Get(env envUtil.Env, appName string) (*models.Application, error) {
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
func IsAliasDefined(env envUtil.Env, appName string, logger *log.Entry) error {
	appAlias := getAlias(env, appName)
	if appAlias != nil {
		logger.Infof("App alias for application %s is defined %s. Now we can try to hit it to see if it responds", appName, *appAlias)
		return nil
	}

	logger.Infof("App alias for application %s is not yet defined", appName)
	return fmt.Errorf("public alias for application %s is not defined", appName)
}

func getAlias(env envUtil.Env, appName string) *string {
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
	return !strings.EqualFold(publicDomainName, canonicalDomainName)
}

// TryGetPublicDomainName Waits for public domain name to be defined
func TryGetPublicDomainName(env envUtil.Env, appName, environmentName, componentName string) (string, error) {
	publicDomainName := getEnvVariable(env, appName, environmentName, componentName, publicDomainNameEnvironmentVariable)
	if publicDomainName == "" {
		return "", fmt.Errorf("public domain name variable for application %s, component %s in environment %s is empty", appName, componentName, environmentName)
	}
	return publicDomainName, nil
}

// TryGetCanonicalDomainName Waits for canonical domain name to be defined
func TryGetCanonicalDomainName(env envUtil.Env, appName, environmentName, componentName string) (string, error) {
	canonicalDomainName := getEnvVariable(env, appName, environmentName, componentName, canonicalEndpointEnvironmentVariable)
	if canonicalDomainName == "" {
		return "", fmt.Errorf("canonical domain name variable for application %s, component %s in environment %s is empty", appName, componentName, environmentName)
	}
	return canonicalDomainName, nil
}

func getEnvVariable(env envUtil.Env, appName, envName, forComponentName, variableName string) string {
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
func AreResponding(env envUtil.Env, logger *log.Entry, urls ...string) error {
	for _, url := range urls {
		responded := IsResponding(env, logger, url)
		if !responded {
			return errors.New("not all endpoints respond")
		}
	}

	return nil
}

// IsResponding Checks if endpoint is responding
func IsResponding(env envUtil.Env, logger *log.Entry, url string) bool {
	req := httpUtils.CreateRequest(env, url, "GET", nil)
	client := http.DefaultClient
	resp, err := client.Do(req)

	if err == nil && resp.StatusCode == 200 {
		logger.Info("App alias responded ok")
		return true
	}

	if err != nil {
		logger.Debugf("Failed request to the alias '%s': %v", url, err)
	}

	if resp != nil {
		logger.Debugf("Request to alias '%s' returned status %v", url, resp.StatusCode)
	}

	if err == nil && resp == nil {
		logger.Debugf("Request to alias '%s' returned no response and no err.", url)
	}

	logger.Infof("Alias '%s' is still not responding", url)
	return false
}
