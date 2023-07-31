package application

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	applicationclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/environment"
	apiclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/platform"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
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
func Register(cfg config.Config, appName, appRepo, appSharedSecret, appCreator, configBranch, configurationItem string, appAdminGroup string, appReaderGroups []string) (*apiclient.RegisterApplicationOK, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()
	bodyParameters := models.ApplicationRegistrationRequest{
		ApplicationRegistration: &models.ApplicationRegistration{
			Name:              &appName,
			Repository:        &appRepo,
			SharedSecret:      &appSharedSecret,
			Creator:           &appCreator,
			AdGroups:          []string{appAdminGroup},
			ReaderAdGroups:    appReaderGroups,
			ConfigBranch:      &configBranch,
			ConfigurationItem: configurationItem,
		},
	}

	params := apiclient.NewRegisterApplicationParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithApplicationRegistration(&bodyParameters)

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetPlatformClient(cfg)

	return client.RegisterApplication(params, clientBearerToken)
}

// DeleteByImpersonatedUser Deletes an application by the impersonated user
func DeleteByImpersonatedUser(cfg config.Config, appName string, logger *log.Entry) error {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()
	logger.Debugf("delete an application %s by the impersonamed user %v, group %s", appName, impersonateUser, impersonateGroup)

	params := applicationclient.NewDeleteApplicationParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithAppName(appName)

	return deleteApplication(cfg, appName, params)
}

// DeleteByServiceAccount an application by the service account
func DeleteByServiceAccount(cfg config.Config, appName string, logger *log.Entry) error {
	err := IsDefined(cfg, appName)
	if err != nil {
		return err
	}
	logger.Debugf("delete an application %s by the service account", appName)

	params := applicationclient.NewDeleteApplicationParams().
		WithAppName(appName)

	return deleteApplication(cfg, appName, params)
}

func RegenerateDeployKey(cfg config.Config, appName, privateKey, sharedSecret string, logger *log.Entry) error {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()
	logger.Debugf("regenerate deploy key for application %s by the impersonamed user %v, group %s", appName, impersonateUser, impersonateGroup)

	params := applicationclient.NewRegenerateDeployKeyParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithAppName(appName).
		WithRegenerateDeployKeyAndSecretData(&models.RegenerateDeployKeyAndSecretData{
			PrivateKey:   privateKey,
			SharedSecret: sharedSecret,
		},
		)

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetApplicationClient(cfg)

	_, err := client.RegenerateDeployKey(params, clientBearerToken)
	if err != nil {
		return fmt.Errorf("failed regenerating deploy key for the application %s: %v", appName, err)
	}
	return nil
}

func HasDeployKey(cfg config.Config, appName, expectedDeployKey string, logger *log.Entry) error {
	actualDeployKey, err := GetDeployKey(cfg, appName, logger)
	if err != nil {
		return err
	}

	if strings.TrimSpace(expectedDeployKey) != strings.TrimSpace(actualDeployKey) {
		return fmt.Errorf("application %s does not have the expected deploy key", appName)
	}

	return nil
}

func IsDeployKeyDefined(cfg config.Config, appName string, logger *log.Entry) error {
	actualDeployKey, err := GetDeployKey(cfg, appName, logger)
	if err != nil {
		return err
	}

	if strings.TrimSpace(actualDeployKey) == "" {
		return fmt.Errorf("deploy key for application %s is not defined", appName)
	}

	return nil
}

func GetDeployKey(cfg config.Config, appName string, logger *log.Entry) (string, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()
	logger.Debugf("get deploy key for application %s by the impersonated user %v, group %s", appName, impersonateUser, impersonateGroup)

	params := applicationclient.NewGetDeployKeyAndSecretParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithAppName(appName)

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetApplicationClient(cfg)

	response, err := client.GetDeployKeyAndSecret(params, clientBearerToken)
	if err != nil {
		return "", fmt.Errorf("failed getting deploy key for the application %s: %v", appName, err)
	}
	return *response.Payload.PublicDeployKey, nil
}

func deleteApplication(cfg config.Config, appName string, params *applicationclient.DeleteApplicationParams) error {
	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetApplicationClient(cfg)

	_, err := client.DeleteApplication(params, clientBearerToken)
	if err != nil {
		return fmt.Errorf("failed deleting the application %s: %v", appName, err)
	}
	return nil
}

// Deploy Deploy application
func Deploy(cfg config.Config, appName, toEnvironment string) (*applicationclient.TriggerPipelineDeployOK, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	bodyParameters := models.PipelineParametersDeploy{
		ToEnvironment: toEnvironment,
	}

	params := applicationclient.NewTriggerPipelineDeployParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithAppName(appName).
		WithPipelineParametersDeploy(&bodyParameters)

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetApplicationClient(cfg)

	return client.TriggerPipelineDeploy(params, clientBearerToken)
}

// IsDefined Checks if application is defined
func IsDefined(cfg config.Config, appName string) error {
	_, err := Get(cfg, appName)
	if err == nil {
		return nil
	}
	return fmt.Errorf("application %s is not defined", appName)
}

func appNamespacesDoNotExist(appName string) error {
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
func DeleteIfExist(cfg config.Config, appName string, logger *log.Entry) error {
	err := DeleteByServiceAccount(cfg, appName, logger)
	if err != nil {
		return nil
	}
	return test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return appNamespacesDoNotExist(appName)
	}, logger)
}

// Get gets an application by appName
func Get(cfg config.Config, appName string) (*models.Application, error) {
	params := applicationclient.NewGetApplicationParams().
		WithAppName(appName).
		WithImpersonateUser(cfg.GetImpersonateUser()).
		WithImpersonateGroup(cfg.GetImpersonateGroups())
	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetApplicationClient(cfg)

	result, err := client.GetApplication(params, clientBearerToken)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}

// IsAliasDefined Checks if app alias is defined
func IsAliasDefined(cfg config.Config, appName string, logger *log.Entry) error {
	appAlias := getAlias(cfg, appName)
	if appAlias != nil {
		logger.Infof("App alias for application %s is defined %s. Now we can try to hit it to see if it responds", appName, *appAlias)
		return nil
	}

	logger.Infof("App alias for application %s is not yet defined", appName)
	return fmt.Errorf("public alias for application %s is not defined", appName)
}

func getAlias(cfg config.Config, appName string) *string {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := applicationclient.NewGetApplicationParams().
		WithAppName(appName).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetApplicationClient(cfg)

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
func TryGetPublicDomainName(cfg config.Config, appName, environmentName, componentName string) (string, error) {
	publicDomainName := getEnvVariable(cfg, appName, environmentName, componentName, publicDomainNameEnvironmentVariable)
	if publicDomainName == "" {
		return "", fmt.Errorf("public domain name variable for application %s, component %s in environment %s is empty", appName, componentName, environmentName)
	}
	return publicDomainName, nil
}

// TryGetCanonicalDomainName Waits for canonical domain name to be defined
func TryGetCanonicalDomainName(cfg config.Config, appName, environmentName, componentName string) (string, error) {
	canonicalDomainName := getEnvVariable(cfg, appName, environmentName, componentName, canonicalEndpointEnvironmentVariable)
	if canonicalDomainName == "" {
		return "", fmt.Errorf("canonical domain name variable for application %s, component %s in environment %s is empty", appName, componentName, environmentName)
	}
	return canonicalDomainName, nil
}

func getEnvVariable(cfg config.Config, appName, envName, forComponentName, variableName string) string {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := environmentclient.NewGetEnvironmentParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetEnvironmentClient(cfg)

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
func AreResponding(logger *log.Entry, urls ...string) error {
	for _, url := range urls {
		responded := IsResponding(logger, url)
		if !responded {
			return errors.New("not all endpoints respond")
		}
	}

	return nil
}

// IsResponding Checks if endpoint is responding
func IsResponding(logger *log.Entry, url string) bool {
	req := httpUtils.CreateRequest(url, "GET", nil)
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
