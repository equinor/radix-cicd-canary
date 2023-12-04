package application

import (
	"context"
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
	errors "github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

const (
	publicDomainNameEnvironmentVariable  = "RADIX_PUBLIC_DOMAIN_NAME"
	canonicalEndpointEnvironmentVariable = "RADIX_CANONICAL_DOMAIN_NAME"
)

// Register Will register application
func Register(ctx context.Context, cfg config.Config, appName, appRepo, appSharedSecret, appCreator, configBranch, configurationItem string, appAdminGroup string, appReaderGroups []string) (*apiclient.RegisterApplicationOK, error) {
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
		WithContext(ctx).
		WithApplicationRegistration(&bodyParameters)

	client := httpUtils.GetPlatformClient(cfg)
	return client.RegisterApplication(params, nil)
}

// DeleteByImpersonatedUser Deletes an application by the impersonated user
func DeleteByImpersonatedUser(ctx context.Context, cfg config.Config, appName string) error {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()
	log.Ctx(ctx).Debug().Msgf("delete an application %s by the impersonamed user %v, group %s", appName, impersonateUser, impersonateGroup)

	params := applicationclient.NewDeleteApplicationParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithContext(ctx).
		WithAppName(appName)

	return deleteApplication(cfg, appName, params)
}

// DeleteByServiceAccount an application by the service account
func DeleteByServiceAccount(ctx context.Context, cfg config.Config, appName string) error {
	err := IsDefined(ctx, cfg, appName)
	if err != nil {
		return err
	}
	log.Ctx(ctx).Debug().Msgf("delete an application by the service account: %s", appName)

	params := applicationclient.NewDeleteApplicationParams().
		WithContext(ctx).
		WithAppName(appName)

	return deleteApplication(cfg, appName, params)
}

func RegenerateDeployKey(ctx context.Context, cfg config.Config, appName, privateKey, sharedSecret string) error {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()
	log.Ctx(ctx).Debug().Str("impersonateGroup", *impersonateGroup).Str("impersonateUser", *impersonateUser).Msg("regenerate deploy key for application by the impersonamed user")

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
		return errors.Wrapf(err, "failed regenerating deploy key for the application %s", appName)
	}
	return nil
}

func HasDeployKey(ctx context.Context, cfg config.Config, appName, expectedDeployKey string) error {
	actualDeployKey, err := GetDeployKey(ctx, cfg, appName)
	if err != nil {
		return err
	}

	if strings.TrimSpace(expectedDeployKey) != strings.TrimSpace(actualDeployKey) {
		return errors.Errorf("application %s does not have the expected deploy key", appName)
	}

	return nil
}

func IsDeployKeyDefined(ctx context.Context, cfg config.Config, appName string) error {
	actualDeployKey, err := GetDeployKey(ctx, cfg, appName)
	if err != nil {
		return err
	}

	if strings.TrimSpace(actualDeployKey) == "" {
		return errors.Errorf("deploy key for application %s is not defined", appName)
	}

	return nil
}

func GetDeployKey(ctx context.Context, cfg config.Config, appName string) (string, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()
	log.Ctx(ctx).Debug().Str("impersonateGroup", *impersonateGroup).Str("impersonateUser", *impersonateUser).Msg("get deploy key for application by the impersonated user")

	params := applicationclient.NewGetDeployKeyAndSecretParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithContext(ctx).
		WithAppName(appName)

	client := httpUtils.GetApplicationClient(cfg)
	response, err := client.GetDeployKeyAndSecret(params, nil)
	if err != nil {
		return "", errors.Wrapf(err, "failed getting deploy key for the application %s", appName)
	}
	return *response.Payload.PublicDeployKey, nil
}

func deleteApplication(cfg config.Config, appName string, params *applicationclient.DeleteApplicationParams) error {
	client := httpUtils.GetApplicationClient(cfg)
	_, err := client.DeleteApplication(params, nil)
	if err != nil {
		return errors.Wrapf(err, "failed deleting the application %s", appName)
	}
	return nil
}

// Deploy Deploy application
func Deploy(ctx context.Context, cfg config.Config, appName, toEnvironment string) (*applicationclient.TriggerPipelineDeployOK, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	bodyParameters := models.PipelineParametersDeploy{
		ToEnvironment: toEnvironment,
	}

	params := applicationclient.NewTriggerPipelineDeployParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithContext(ctx).
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
	return errors.Errorf("application %s is not defined", appName)
}

func appNamespacesDoNotExist(ctx context.Context, appName string) error {
	nsList, err := kubeUtils.GetKubernetesClient().CoreV1().Namespaces().List(ctx, metav1.ListOptions{
		LabelSelector: labels.Set{"radix-app": appName}.String(),
	})
	if err != nil {
		return errors.WithStack(err)
	}
	if len(nsList.Items) > 0 {
		return errors.Errorf("there are %d namespaces for the application %s", len(nsList.Items), appName)
	}
	return nil
}

// DeleteIfExist Delete application if it exists
func DeleteIfExist(ctx context.Context, cfg config.Config, appName string) error {
	err := DeleteByServiceAccount(ctx, cfg, appName)
	if err != nil {
		return nil
	}
	return test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		return appNamespacesDoNotExist(ctx, appName)
	})
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
		return nil, errors.WithStack(err)
	}
	return result.Payload, nil
}

// IsAliasDefined Checks if app alias is defined
func IsAliasDefined(ctx context.Context, cfg config.Config, appName string) error {
	appAlias := getAlias(ctx, cfg, appName)
	if appAlias != nil {
		log.Ctx(ctx).Info().Msgf("App alias for application %s is defined: %s", appName, *appAlias)
		return nil
	}

	log.Ctx(ctx).Info().Msg("App alias for application is not yet defined")
	return errors.Errorf("public alias for application %s is not defined", appName)
}

func getAlias(ctx context.Context, cfg config.Config, appName string) *string {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := applicationclient.NewGetApplicationParams().
		WithAppName(appName).
		WithContext(ctx).
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
func TryGetPublicDomainName(ctx context.Context, cfg config.Config, appName, environmentName, componentName string) (string, error) {
	publicDomainName := getEnvVariable(ctx, cfg, appName, environmentName, componentName, publicDomainNameEnvironmentVariable)
	if publicDomainName == "" {
		return "", errors.Errorf("public domain name variable for application %s, component %s in environment %s is empty", appName, componentName, environmentName)
	}
	return publicDomainName, nil
}

// TryGetCanonicalDomainName Waits for canonical domain name to be defined
func TryGetCanonicalDomainName(ctx context.Context, cfg config.Config, appName, environmentName, componentName string) (string, error) {
	canonicalDomainName := getEnvVariable(ctx, cfg, appName, environmentName, componentName, canonicalEndpointEnvironmentVariable)
	if canonicalDomainName == "" {
		return "", errors.Errorf("canonical domain name variable for application %s, component %s in environment %s is empty", appName, componentName, environmentName)
	}
	return canonicalDomainName, nil
}

func getEnvVariable(ctx context.Context, cfg config.Config, appName, envName, forComponentName, variableName string) string {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := environmentclient.NewGetEnvironmentParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithContext(ctx).
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
		logger.Debug().Msgf("Failed request to %s with the alias: %v", url, err)
	}

	if resp != nil {
		logger.Debug().Msgf("Request to alias '%s' returned status %v", url, resp.StatusCode)
	}

	if err == nil && resp == nil {
		logger.Debug().Msgf("Request to alias returned, no response and no err: %s", url)
	}

	logger.Info().Msgf("Alias is still not responding: %s", url)
	return false
}
