package adgroup

import (
	"errors"
	"fmt"

	apiclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/environment"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/go-openapi/runtime"
	log "github.com/sirupsen/logrus"
)

type step struct {
	logger *log.Entry
}

const (
	adGroupWithNoAccess = "12345678-9012-3456-7890-123456789012"
)

// Update Tests that updates to AD group locks down an application
func Update(cfg config.Config, suiteName string) error {
	s := &step{logger: log.WithFields(log.Fields{"Suite": suiteName})}

	s.logger.Debugf("check that admin AD-Group has access")
	err := test.WaitForCheckFuncOrTimeout(cfg, s.hasAccess, s.logger)
	if err != nil {
		return fmt.Errorf("failed to get update details of the suite %s: %w", suiteName, err)
	}
	s.logger.Debugf("admin AD-Group has access")

	s.logger.Debugf("patch the RR and set new admin AD group, which cicd-canary does not impersonate")
	err = patchAdGroup(cfg, adGroupWithNoAccess)
	if err != nil {
		return err
	}
	s.logger.Debugf("RR's admin AD-Group is patched")

	s.logger.Debugf("check that cicd-canary lacks access with current impersonation")
	err = test.WaitForCheckFuncOrTimeout(cfg, s.hasNoAccess, s.logger)
	if err != nil {
		return fmt.Errorf("failed to get patchAdGroup update details: %w", err)
	}
	s.logger.Debugf("cicd-canary lacks access with current group impersonation")

	s.logger.Debugf("patch the RR and set original AD group as admin, which cicd-canary does impersonate")
	err = patchAdGroup(cfg, cfg.GetAppAdminGroup())
	if err != nil {
		return err
	}
	s.logger.Debugf("admin AD-Group is patched")

	s.logger.Debugf("check that cicd-canary has access")
	err = test.WaitForCheckFuncOrTimeout(cfg, s.hasAccess, s.logger)
	s.logger.Debugf("cicd-canary has access")
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

	s.logger.Debugf(" - accessToApplication: %v, accessToBuild: %v, accessToSecret: %v", accessToApplication, accessToBuild, accessToSecret)

	hasProperAccess := accessToApplication == properAccess && accessToBuild == properAccess && accessToSecret == properAccess
	s.logger.Debugf(" - hasProperAccess: %v", hasProperAccess)

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

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetApplicationClient(cfg)

	_, err := client.ModifyRegistrationDetails(params, clientBearerToken)
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

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetApplicationClient(cfg)

	application, err := client.GetApplication(params, clientBearerToken)
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

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetApplicationClient(cfg)

	_, err := client.TriggerPipelineBuild(params, clientBearerToken)
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
				SecretValue: stringPtr(defaults.App2SecretValue),
			})

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetEnvironmentClient(cfg)

	_, err := client.ChangeComponentSecret(params, clientBearerToken)
	if err != nil {
		return fmt.Errorf("error calling ChangeComponentSecret for application %s: %w", defaults.App2Name, err)
	}
	return nil
}

func (s *step) isChangeComponentSecretForbidden(err error) bool {
	if errors.Is(err, &environmentclient.ChangeComponentSecretForbidden{}) {
		return true
	}
	s.logger.Debugf("ChangeComponentSecret err: %v", err)
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
		s.logger.Debugf("checkErrorResponse err code: %d", 403)
		return true
	case *runtime.APIError:
		s.logger.Debugf("checkErrorResponse err code: %d", err.Code)
		return err.Code == expectedStatusCode
	default:
		s.logger.Debugf("checkErrorResponse err is not an expected type")
		return false
	}
}

func stringPtr(str string) *string {
	return &str
}
