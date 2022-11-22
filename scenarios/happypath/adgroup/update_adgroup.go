package adgroup

import (
	"errors"
	"fmt"
	apiclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/client/environment"
	"github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
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
func Update(env envUtil.Env, suiteName string) error {
	s := &step{logger: log.WithFields(log.Fields{"Suite": suiteName})}

	s.logger.Debugf("check that admin AD-Group has access")
	err := test.WaitForCheckFuncOrTimeout(env, s.hasAccess, s.logger)
	if err != nil {
		return fmt.Errorf("failed to get update details of the suite %s: %w", suiteName, err)
	}
	s.logger.Debugf("admin AD-Group has access")

	s.logger.Debugf("patch an admin AD-Group without access")
	err = patchAdGroup(env, adGroupWithNoAccess)
	if err != nil {
		return err
	}
	s.logger.Debugf("admin AD-Group is patched")

	s.logger.Debugf("check that admin AD-Group has no access")
	err = test.WaitForCheckFuncOrTimeout(env, s.hasNoAccess, s.logger)
	if err != nil {
		return fmt.Errorf("failed to get patchAdGroup update details: %w", err)
	}
	s.logger.Debugf("admin AD-Group has no access")

	s.logger.Debugf("patch an admin AD-Group with access")
	err = patchAdGroup(env, env.GetImpersonateGroup())
	if err != nil {
		return err
	}
	s.logger.Debugf("admin AD-Group is patched")

	s.logger.Debugf("check that admin AD-Group has access")
	err = test.WaitForCheckFuncOrTimeout(env, s.hasAccess, s.logger)
	s.logger.Debugf("admin AD-Group has no access")
	return err
}

func (s *step) hasNoAccess(env envUtil.Env) error {
	return s.hasProperAccess(env, false)
}

func (s *step) hasAccess(env envUtil.Env) error {
	return s.hasProperAccess(env, true)
}

func (s *step) hasProperAccess(env envUtil.Env, properAccess bool) error {
	_, err := getApplication(env)
	accessToApplication := !isGetApplicationForbidden(err)

	err = buildApp(env)
	accessToBuild := !s.isTriggerPipelineBuildForbidden(err)

	err = setSecret(env)
	accessToSecret := !s.isChangeComponentSecretForbidden(err)

	s.logger.Debugf(" - accessToApplication: %v, accessToBuild: %v, accessToSecret: %v", accessToApplication, accessToBuild, accessToSecret)

	hasProperAccess := accessToApplication == properAccess && accessToBuild == properAccess && accessToSecret == properAccess
	s.logger.Debugf(" - hasProperAccess: %v", hasProperAccess)

	if !hasProperAccess {
		return fmt.Errorf("proper access hasn't been granted yet")
	}
	return nil
}

func patchAdGroup(env envUtil.Env, adGroup string) error {
	patchRequest := models.ApplicationRegistrationPatchRequest{
		ApplicationRegistrationPatch: &models.ApplicationRegistrationPatch{
			AdGroups: []string{adGroup},
		},
	}

	params := apiclient.NewModifyRegistrationDetailsParams().
		WithAppName(config.App2Name).
		WithPatchRequest(&patchRequest)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	_, err := client.ModifyRegistrationDetails(params, clientBearerToken)
	if err != nil {
		return err
	}

	return nil
}

func getApplication(env envUtil.Env) (*models.Application, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := apiclient.NewGetApplicationParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(config.App2Name)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	application, err := client.GetApplication(params, clientBearerToken)
	if err != nil {
		return nil, err
	}

	return application.Payload, nil
}

func buildApp(env envUtil.Env) error {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	bodyParameters := models.PipelineParametersBuild{
		Branch: config.App2BranchToBuildFrom,
	}

	params := apiclient.NewTriggerPipelineBuildParams().
		WithAppName(config.App2Name).
		WithPipelineParametersBuild(&bodyParameters).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	_, err := client.TriggerPipelineBuild(params, clientBearerToken)
	if err != nil {
		return err
	}

	return nil
}

func setSecret(env envUtil.Env) error {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := environmentclient.NewChangeComponentSecretParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(config.App2Name).
		WithEnvName(config.App2EnvironmentName).
		WithComponentName(config.App2Component2Name).
		WithSecretName(config.App2SecretName).
		WithComponentSecret(
			&models.SecretParameters{
				SecretValue: stringPtr(config.App2SecretValue),
			})

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetEnvironmentClient(env)

	_, err := client.ChangeComponentSecret(params, clientBearerToken)
	if err != nil {
		return fmt.Errorf("error calling ChangeComponentSecret for application %s: %w", config.App2Name, err)
	}
	return nil
}

func (s *step) isChangeComponentSecretForbidden(err error) bool {
	if errors.Is(err, &environmentclient.ChangeComponentSecretForbidden{}) {
		return true
	}
	s.logger.Debugf("ChangeComponentSecret error: %v", err)
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
	apiError, ok := err.(*runtime.APIError)
	if ok {
		errorCode := apiError.Code
		s.logger.Debugf("checkErrorResponse error code: %d", errorCode)
		if errorCode == expectedStatusCode {
			return true
		}
	} else {
		s.logger.Debugf("checkErrorResponse error is not runtime.APIError")
	}
	return false
}

func stringPtr(str string) *string {
	return &str
}
