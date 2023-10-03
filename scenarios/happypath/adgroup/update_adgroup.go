package adgroup

import (
	"context"
	"errors"
	"fmt"

	apiclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/environment"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	commonUtils "github.com/equinor/radix-common/utils"
	"github.com/go-openapi/runtime"
	"github.com/rs/zerolog/log"
)

const (
	adGroupWithNoAccess = "12345678-9012-3456-7890-123456789012"
)

// Update Tests that updates to AD group locks down an application
func Update(ctx context.Context, cfg config.Config) error {
	logger := log.Ctx(ctx)

	logger.Debug().Msg("check that admin AD-Group has access")
	err := test.WaitForCheckFuncOrTimeout(ctx, cfg, hasAccess)
	if err != nil {
		return fmt.Errorf("failed to get update details of the suite: %w", err)
	}
	logger.Debug().Msg("admin AD-Group has access")

	logger.Debug().Msg("patch the RR and set new admin AD group, which the impersonated user is not member of")
	err = patchAdGroup(ctx, cfg, defaults.App2Name, adGroupWithNoAccess)
	if err != nil {
		return err
	}
	logger.Debug().Msg("RR's admin AD-Group is patched")

	logger.Debug().Msg("check that the application cannot be accessed with current impersonation")
	err = test.WaitForCheckFuncOrTimeout(ctx, cfg, hasNoAccess)
	if err != nil {
		return fmt.Errorf("failed to get patchAdGroup update details: %w", err)
	}
	logger.Debug().Msg("application cannot be accessed with current impersonation")

	logger.Debug().Msg("patch the RR and set oroginal admin AD group, which the impersonated user is member of")
	err = patchAdGroup(ctx, cfg, defaults.App2Name, cfg.GetAppAdminGroup())
	if err != nil {
		return err
	}
	logger.Debug().Msg("admin AD-Group is patched")

	logger.Debug().Msg("check that the application can be accessed with current impersonation")
	err = test.WaitForCheckFuncOrTimeout(ctx, cfg, hasAccess)
	logger.Debug().Msg("application can be accessed with current impersonation")
	return err
}

func hasNoAccess(cfg config.Config, ctx context.Context) error {
	return hasProperAccess(ctx, cfg, false)
}

func hasAccess(cfg config.Config, ctx context.Context) error {
	return hasProperAccess(ctx, cfg, true)
}

func hasProperAccess(ctx context.Context, cfg config.Config, properAccess bool) error {
	_, err := getApplication(ctx, cfg)
	accessToApplication := !isGetApplicationForbidden(err)

	err = buildApp(cfg)
	accessToBuild := !isTriggerPipelineBuildForbidden(ctx, err)

	err = setSecret(cfg)
	accessToSecret := !isChangeComponentSecretForbidden(ctx, err)
	hasProperAccess := accessToApplication == properAccess && accessToBuild == properAccess && accessToSecret == properAccess

	log.Ctx(ctx).Debug().Msgf("AccessToApplication: %v, accessToBuild: %v, accessToSecret: %v, HasProperAccess: %v", accessToApplication, accessToBuild, accessToSecret, hasProperAccess)

	if !hasProperAccess {
		return fmt.Errorf("proper access hasn't been granted yet")
	}
	return nil
}

func patchAdGroup(ctx context.Context, cfg config.Config, appName string, adGroup string) error {
	patchRequest := models.ApplicationRegistrationPatchRequest{
		ApplicationRegistrationPatch: &models.ApplicationRegistrationPatch{
			AdGroups: []string{adGroup},
		},
	}

	params := apiclient.NewModifyRegistrationDetailsParams().
		WithAppName(appName).
		WithContext(ctx).
		WithPatchRequest(&patchRequest)

	client := httpUtils.GetApplicationClient(cfg)
	_, err := client.ModifyRegistrationDetails(params, nil)
	if err != nil {
		return err
	}

	return nil
}

func getApplication(ctx context.Context, cfg config.Config) (*models.Application, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := apiclient.NewGetApplicationParams().
		WithImpersonateUser(impersonateUser).
		WithContext(ctx).
		WithImpersonateGroup(impersonateGroup).
		WithAppName(defaults.App2Name)

	client := httpUtils.GetApplicationClient(cfg)
	application, err := client.GetApplication(params, nil)
	if err != nil {
		return nil, err
	}

	return application.Payload, nil
}

func buildApp(cfg config.Config) error {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	bodyParameters := models.PipelineParametersBuild{
		Branch: defaults.App2BranchToBuildFrom,
	}

	params := apiclient.NewTriggerPipelineBuildParams().
		WithAppName(defaults.App2Name).
		WithPipelineParametersBuild(&bodyParameters).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	client := httpUtils.GetApplicationClient(cfg)
	_, err := client.TriggerPipelineBuild(params, nil)
	if err != nil {
		return err
	}

	return nil
}

func setSecret(cfg config.Config) error {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := environmentclient.NewChangeComponentSecretParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithAppName(defaults.App2Name).
		WithEnvName(defaults.App2EnvironmentName).
		WithComponentName(defaults.App2Component2Name).
		WithSecretName(defaults.App2SecretName).
		WithComponentSecret(
			&models.SecretParameters{
				SecretValue: commonUtils.StringPtr(defaults.App2SecretValue),
			})

	client := httpUtils.GetEnvironmentClient(cfg)
	_, err := client.ChangeComponentSecret(params, nil)
	if err != nil {
		return fmt.Errorf("error calling ChangeComponentSecret for application %s: %w", defaults.App2Name, err)
	}
	return nil
}

func isChangeComponentSecretForbidden(ctx context.Context, err error) bool {
	if errors.Is(err, &environmentclient.ChangeComponentSecretForbidden{}) {
		return true
	}
	log.Ctx(ctx).Debug().Msgf("ChangeComponentSecret: %v", err)
	return false
}

func isGetApplicationForbidden(err error) bool {
	switch err.(type) {
	case *apiclient.GetApplicationForbidden:
		return true
	}

	return false
}

func isTriggerPipelineBuildForbidden(ctx context.Context, err error) bool {
	return err != nil && checkErrorResponse(ctx, err, 403)
}

func checkErrorResponse(ctx context.Context, err error, expectedStatusCode int) bool {
	switch err := err.(type) {
	case *apiclient.TriggerPipelineBuildForbidden:
		log.Ctx(ctx).Debug().Int("errorCode", 403).Msg("checkErrorResponse err code")
		return true
	case *runtime.APIError:
		log.Ctx(ctx).Debug().Int("errorCode", err.Code).Msg("checkErrorResponse err code")
		return err.Code == expectedStatusCode
	default:
		log.Ctx(ctx).Debug().Msg("checkErrorResponse err is not an expected type")
		return false
	}
}
