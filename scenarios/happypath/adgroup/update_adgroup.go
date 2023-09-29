package adgroup

import (
	"errors"
	"fmt"

	commonUtils "github.com/equinor/radix-common/utils"
	"github.com/rs/zerolog"

	apiclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/environment"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/go-openapi/runtime"
	"github.com/rs/zerolog/log"
)

type step struct {
	logger zerolog.Logger
}

const (
	adGroupWithNoAccess = "12345678-9012-3456-7890-123456789012"
)

// Update Tests that updates to AD group locks down an application
func Update(cfg config.Config, suiteName string) error {
	s := &step{logger: log.With().Str("suite", suiteName).Logger()}

	s.logger.Debug().Msg("check that admin AD-Group has access")
	err := test.WaitForCheckFuncOrTimeout(cfg, s.hasAccess, s.logger)
	if err != nil {
		return fmt.Errorf("failed to get update details of the suite %s: %w", suiteName, err)
	}
	s.logger.Debug().Msg("admin AD-Group has access")

	s.logger.Debug().Msg("patch the RR and set new admin AD group, which the impersonated user is not member of")
	err = patchAdGroup(cfg, adGroupWithNoAccess)
	if err != nil {
		return err
	}
	s.logger.Debug().Msg("RR's admin AD-Group is patched")

	s.logger.Debug().Msg("check that the application cannot be accessed with current impersonation")
	err = test.WaitForCheckFuncOrTimeout(cfg, s.hasNoAccess, s.logger)
	if err != nil {
		return fmt.Errorf("failed to get patchAdGroup update details: %w", err)
	}
	s.logger.Debug().Msg("application cannot be accessed with current impersonation")

	s.logger.Debug().Msg("patch the RR and set oroginal admin AD group, which the impersonated user is member of")
	err = patchAdGroup(cfg, cfg.GetAppAdminGroup())
	if err != nil {
		return err
	}
	s.logger.Debug().Msg("admin AD-Group is patched")

	s.logger.Debug().Msg("check that the application can be accessed with current impersonation")
	err = test.WaitForCheckFuncOrTimeout(cfg, s.hasAccess, s.logger)
	s.logger.Debug().Msg("application can be accessed with current impersonation")
	return err
}

func (s *step) hasNoAccess(cfg config.Config) error {
	return s.hasProperAccess(cfg, false)
}

func (s *step) hasAccess(cfg config.Config) error {
	return s.hasProperAccess(cfg, true)
}

func (s *step) hasProperAccess(cfg config.Config, properAccess bool) error {
	_, err := getApplication(cfg)
	accessToApplication := !isGetApplicationForbidden(err)

	err = buildApp(cfg)
	accessToBuild := !s.isTriggerPipelineBuildForbidden(err)

	err = setSecret(cfg)
	accessToSecret := !s.isChangeComponentSecretForbidden(err)

	s.logger.Debug().Msgf(" - accessToApplication: %v, accessToBuild: %v, accessToSecret: %v", accessToApplication, accessToBuild, accessToSecret)

	hasProperAccess := accessToApplication == properAccess && accessToBuild == properAccess && accessToSecret == properAccess
	s.logger.Debug().Msgf(" - hasProperAccess: %v", hasProperAccess)

	if !hasProperAccess {
		return fmt.Errorf("proper access hasn't been granted yet")
	}
	return nil
}

func patchAdGroup(cfg config.Config, adGroup string) error {
	patchRequest := models.ApplicationRegistrationPatchRequest{
		ApplicationRegistrationPatch: &models.ApplicationRegistrationPatch{
			AdGroups: []string{adGroup},
		},
	}

	params := apiclient.NewModifyRegistrationDetailsParams().
		WithAppName(defaults.App2Name).
		WithPatchRequest(&patchRequest)

	client := httpUtils.GetApplicationClient(cfg)
	_, err := client.ModifyRegistrationDetails(params, nil)
	if err != nil {
		return err
	}

	return nil
}

func getApplication(cfg config.Config) (*models.Application, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := apiclient.NewGetApplicationParams().
		WithImpersonateUser(impersonateUser).
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

func (s *step) isChangeComponentSecretForbidden(err error) bool {
	if errors.Is(err, &environmentclient.ChangeComponentSecretForbidden{}) {
		return true
	}
	s.logger.Debug().Err(err).Msg("ChangeComponentSecret")
	return false
}

func isGetApplicationForbidden(err error) bool {
	switch err.(type) {
	case *apiclient.GetApplicationForbidden:
		return true
	}

	return false
}

func (s *step) isTriggerPipelineBuildForbidden(err error) bool {
	return err != nil && s.checkErrorResponse(err, 403)
}

func (s *step) checkErrorResponse(err error, expectedStatusCode int) bool {
	switch err := err.(type) {
	case *apiclient.TriggerPipelineBuildForbidden:
		s.logger.Debug().Int("errorCode", 403).Msg("checkErrorResponse err code")
		return true
	case *runtime.APIError:
		s.logger.Debug().Int("errorCode", err.Code).Msg("checkErrorResponse err code")
		return err.Code == expectedStatusCode
	default:
		s.logger.Debug().Msg("checkErrorResponse err is not an expected type")
		return false
	}
}
