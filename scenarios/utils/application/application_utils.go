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
	"github.com/rs/zerolog/log"
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

	client := httpUtils.GetPlatformClient(cfg)
	return client.RegisterApplication(params, nil)
}

// DeleteByImpersonatedUser Deletes an application by the impersonated user
func DeleteByImpersonatedUser(cfg config.Config, appName string, ctx context.Context) error {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()
	log.Ctx(ctx).Debug().Msgf("delete an application %s by the impersonamed user %v, group %s", appName, impersonateUser, impersonateGroup)

	params := applicationclient.NewDeleteApplicationParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithAppName(appName)

	return deleteApplication(cfg, appName, params)
}

// DeleteByServiceAccount an application by the service account
func DeleteByServiceAccount(cfg config.Config, appName string, ctx context.Context) error {
	err := IsDefined(ctx, cfg, appName)
	if err != nil {
		return err
	}
	log.Ctx(ctx).Debug().Msg("delete an application by the service account")

	params := applicationclient.NewDeleteApplicationParams().
		WithAppName(appName)

	return deleteApplication(cfg, appName, params)
}

func RegenerateDeployKey(cfg config.Config, appName, privateKey, sharedSecret string, ctx context.Context) error {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()
	log.Ctx(ctx).Debug().Strs("impersonateGroup", impersonateGroup).Str("impersonateUser", *impersonateUser).Msg("regenerate deploy key for application by the impersonamed user")

	params := applicationclient.NewRegenerateDeployKeyParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithAppName(appName).
		WithRegenerateDeployKeyAndSecretData(&models.RegenerateDeployKeyAndSecretData{
			PrivateKey:   privateKey,
			SharedSecret: sharedSecret,
		},
		)

	client := httpUtils.GetApplicationClient(cfg)
	_, err := client.RegenerateDeployKey(params, nil)
	if err != nil {
		return fmt.Errorf("failed regenerating deploy key for the application %s: %v", appName, err)
	}
	return nil
}

func HasDeployKey(cfg config.Config, appName, expectedDeployKey string, ctx context.Context) error {
	actualDeployKey, err := GetDeployKey(cfg, appName, ctx)
	if err != nil {
		return err
	}

	if strings.TrimSpace(expectedDeployKey) != strings.TrimSpace(actualDeployKey) {
		return fmt.Errorf("application %s does not have the expected deploy key", appName)
	}

	return nil
}

func IsDeployKeyDefined(cfg config.Config, appName string, ctx context.Context) error {
	actualDeployKey, err := GetDeployKey(cfg, appName, ctx)
	if err != nil {
		return err
	}

	if strings.TrimSpace(actualDeployKey) == "" {
		return fmt.Errorf("deploy key for application %s is not defined", appName)
	}

	return nil
}

func GetDeployKey(cfg config.Config, appName string, ctx context.Context) (string, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()
	log.Ctx(ctx).Debug().Strs("impersonateGroup", impersonateGroup).Str("impersonateUser", *impersonateUser).Msg("get deploy key for application by the impersonated user")

	params := applicationclient.NewGetDeployKeyAndSecretParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithAppName(appName)

	client := httpUtils.GetApplicationClient(cfg)
	response, err := client.GetDeployKeyAndSecret(params, nil)
	if err != nil {
		return "", fmt.Errorf("failed getting deploy key for the application %s: %v", appName, err)
	}
	return *response.Payload.PublicDeployKey, nil
}

func deleteApplication(cfg config.Config, appName string, params *applicationclient.DeleteApplicationParams) error {
	client := httpUtils.GetApplicationClient(cfg)
	_, err := client.DeleteApplication(params, nil)
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

	client := httpUtils.GetApplicationClient(cfg)
	return client.TriggerPipelineDeploy(params, nil)
}

// IsDefined Checks if application is defined
func IsDefined(ctx context.Context, cfg config.Config, appName string) error {
	_, err := Get(ctx, cfg, appName)
	if err == nil {
		return nil
	}
	return fmt.Errorf("application %s is not defined", appName)
}

func appNamespacesDoNotExist(ctx context.Context, appName string) error {
	nsList, err := kubeUtils.GetKubernetesClient().CoreV1().Namespaces().List(ctx, metav1.ListOptions{
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
func DeleteIfExist(cfg config.Config, appName string, ctx context.Context) error {
	err := DeleteByServiceAccount(cfg, appName, ctx)
	if err != nil {
		return nil
	}
	return test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config, ctx context.Context) error {
		return appNamespacesDoNotExist(ctx, appName)
	}, ctx)
}

// Get gets an application by appName
func Get(ctx context.Context, cfg config.Config, appName string) (*models.Application, error) {
	params := applicationclient.NewGetApplicationParams().
		WithAppName(appName).
		WithContext(ctx).
		WithImpersonateUser(cfg.GetImpersonateUser()).
		WithImpersonateGroup(cfg.GetImpersonateGroups())
	client := httpUtils.GetApplicationClient(cfg)
	result, err := client.GetApplication(params, nil)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}

// IsAliasDefined Checks if app alias is defined
func IsAliasDefined(cfg config.Config, appName string, ctx context.Context) error {
	appAlias := getAlias(cfg, appName)
	if appAlias != nil {
		log.Ctx(ctx).Info().Str("appAlias", *appAlias).Msg("App alias for application is defined. Now we can try to hit it to see if it responds")
		return nil
	}

	log.Ctx(ctx).Info().Msg("App alias for application is not yet defined")
	return fmt.Errorf("public alias for application %s is not defined", appName)
}

func getAlias(cfg config.Config, appName string) *string {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := applicationclient.NewGetApplicationParams().
		WithAppName(appName).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)
	client := httpUtils.GetApplicationClient(cfg)
	applicationDetails, err := client.GetApplication(params, nil)
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
	client := httpUtils.GetEnvironmentClient(cfg)
	environmentDetails, err := client.GetEnvironment(params, nil)
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
func AreResponding(ctx context.Context, urls ...string) error {
	for _, url := range urls {
		responded := IsResponding(ctx, url)
		if !responded {
			return errors.New("not all endpoints respond")
		}
	}

	return nil
}

// IsResponding Checks if endpoint is responding
func IsResponding(ctx context.Context, url string) bool {
	req := httpUtils.CreateRequest(url, "GET", nil)
	client := http.DefaultClient
	resp, err := client.Do(req)
	logger := log.Ctx(ctx)

	if err == nil && resp.StatusCode == 200 {
		logger.Info().Msg("App alias responded ok")
		return true
	}

	if err != nil {
		logger.Debug().Err(err).Str("url", url).Msg("Failed request to the alias")
	}

	if resp != nil {
		logger.Debug().Msgf("Request to alias '%s' returned status %v", url, resp.StatusCode)
	}

	if err == nil && resp == nil {
		logger.Debug().Str("url", url).Msg("Request to alias returned, no response and no err")
	}

	logger.Info().Str("url", url).Msgf("Alias is still not responding")
	return false
}
