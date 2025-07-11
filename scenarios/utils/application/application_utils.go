package application

import (
	"context"
	"strings"

	applicationclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
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
	log.Ctx(ctx).Debug().Msgf("delete an application %s by the impersonamed user %v, group %s", appName, impersonateUser, *impersonateGroup)

	params := applicationclient.NewDeleteApplicationParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithContext(ctx).
		WithAppName(appName)

	return deleteApplication(cfg, appName, params)
}

// DeleteByServiceAccount an application by the service account
func DeleteByServiceAccount(ctx context.Context, cfg config.Config, appName string) error {
	log.Ctx(ctx).Debug().Msgf("delete an application by the service account: %s", appName)

	params := applicationclient.NewDeleteApplicationParams().
		WithContext(ctx).
		WithAppName(appName)

	return deleteApplication(cfg, appName, params)
}

func RegenerateDeployKey(ctx context.Context, cfg config.Config, appName, privateKey string) error {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()
	log.Ctx(ctx).Debug().Str("impersonateGroup", *impersonateGroup).Str("impersonateUser", *impersonateUser).Msg("regenerate deploy key for application by the impersonamed user")

	params := applicationclient.NewRegenerateDeployKeyParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithAppName(appName).
		WithRegenerateDeployKeyAndSecretData(&models.RegenerateDeployKeyData{PrivateKey: privateKey})

	client := httpUtils.GetApplicationClient(cfg)
	_, err := client.RegenerateDeployKey(params, nil)
	return err
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
        if _, ok := err.(*applicationclient.DeleteApplicationNotFound); ok {
            return nil
        }
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
