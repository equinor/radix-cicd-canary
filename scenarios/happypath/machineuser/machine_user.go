package machineuser

import (
	"errors"
	"fmt"

	apiclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/client/environment"
	"github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	log "github.com/sirupsen/logrus"
)

type step struct {
	logger *log.Entry
}

// Create Tests that machine user is created properly
func Create(cfg config.Config, suiteName string) error {
	s := &step{logger: log.WithFields(log.Fields{"Suite": suiteName})}

	s.logger.Debugf("patch to enable machine user")
	const enabled = true
	err := s.patchMachineUser(cfg, enabled)
	if err != nil {
		return err
	}
	s.logger.Debugf("patches")

	s.logger.Debugf("get machine user token")
	machineUserToken, err := test.WaitForCheckFuncWithValueOrTimeout(cfg, func(cfg config.Config) (*string, error) {
		return getMachineUserToken(cfg)
	}, s.logger)

	if err != nil {
		return err
	}
	s.logger.Debugf("machine user token is given")

	s.logger.Debugf("check machine user token has access")
	err = test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return s.hasAccess(cfg, *machineUserToken)
	}, s.logger)

	if err != nil {
		return fmt.Errorf("does not have expected access with machine token. %w", err)
	}

	// Should only have access to its own application
	hasAccessToOtherApplication := s.hasAccessToApplication(cfg, defaults.App1Name, *machineUserToken)
	if hasAccessToOtherApplication {
		return fmt.Errorf("has not expected access to another application '%s'", defaults.App1Name)
	}

	// Disable machine user
	err = s.patchMachineUser(cfg, !enabled)
	if err != nil {
		return err
	}

	// Token should no longer have access
	s.logger.Debugf("check machine user token has no access")
	err = test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return s.hasNoAccess(cfg, *machineUserToken)
	}, s.logger)

	if err != nil {
		return fmt.Errorf("has not expected access with machine token. %w", err)
	}
	s.logger.Debug("MachineUser was set and un-set properly")
	return nil
}

func getMachineUserToken(cfg config.Config) (*string, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroup()

	params := apiclient.NewRegenerateMachineUserTokenParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(defaults.App2Name)

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetApplicationClient(cfg)

	tokenResponse, err := client.RegenerateMachineUserToken(params, clientBearerToken)
	if err != nil {
		return nil, fmt.Errorf("cannot regenerate machine token: %v", err)
	}

	return tokenResponse.Payload.Token, nil
}

func (s *step) patchMachineUser(cfg config.Config, enabled bool) error {
	s.logger.Debugf("Set MachineUser to %v", enabled)
	patchRequest := models.ApplicationRegistrationPatchRequest{
		ApplicationRegistrationPatch: &models.ApplicationRegistrationPatch{
			MachineUser: &enabled,
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
	s.logger.Debugf("MachineUser has been set to %v", enabled)
	return nil
}

func (s *step) hasNoAccess(cfg config.Config, machineUserToken string) error {
	return s.hasProperAccess(cfg, machineUserToken, false)
}

func (s *step) hasAccess(cfg config.Config, machineUserToken string) error {
	return s.hasProperAccess(cfg, machineUserToken, true)
}

func (s *step) hasProperAccess(cfg config.Config, machineUserToken string, properAccess bool) error {
	s.logger.Debugf("check hasProperAccess: %v", properAccess)
	accessToApplication := s.hasAccessToApplication(cfg, defaults.App2Name, machineUserToken)

	err := buildApp(cfg, machineUserToken)
	s.logger.Debugf("err from buildApp: %v", err)
	accessToBuild := !s.isTriggerPipelineBuildUnauthorized(err)

	err = setSecret(cfg, machineUserToken)
	s.logger.Debugf("err from setSecret: %v", err)
	accessToSecret := !s.isSetSecretUnauthorizedError(err)

	hasProperAccess := accessToApplication == properAccess && accessToBuild == properAccess && accessToSecret == properAccess
	s.logger.Debugf(" - accessToApplication: %v, accessToBuild: %v, accessToSecret: %v", accessToApplication, accessToBuild, accessToSecret)
	s.logger.Debugf(" - hasProperAccess: %v", hasProperAccess)
	if !hasProperAccess {
		return fmt.Errorf("proper access hasn't been granted yet")
	}
	return nil
}

func (s *step) hasAccessToApplication(cfg config.Config, appName, machineUserToken string) bool {
	s.logger.Debugf("get application with machine user token")
	_, err := getApplication(cfg, appName, machineUserToken)
	if err != nil {
		s.logger.Debugf("got an err %v", err)
	}
	isGetApplicationUnauthorized := s.isGetApplicationUnauthorized(err)
	getApplicationForbidden := s.isGetApplicationForbidden(err)
	s.logger.Debugf("isGetApplicationUnauthorized: %v, getApplicationForbidden: %v", isGetApplicationUnauthorized, getApplicationForbidden)
	hasAccess := !isGetApplicationUnauthorized && !getApplicationForbidden
	s.logger.Debugf("hasAccessToApplication: %v", hasAccess)
	return hasAccess
}

func (s *step) isGetApplicationUnauthorized(err error) bool {
	if errors.Is(err, &apiclient.GetApplicationUnauthorized{}) {
		return true
	}
	s.logger.Debugf("GetApplicationUnauthorized err: %v", err)
	return false
}

func (s *step) isSetSecretUnauthorizedError(err error) bool {
	if errors.Is(err, &environmentclient.ChangeComponentSecretUnauthorized{}) {
		return true
	}
	s.logger.Debugf("SetSecretUnauthorized err: %v", err)
	return false
}

func getApplication(cfg config.Config, appName, machineUserToken string) (*models.Application, error) {
	params := apiclient.NewGetApplicationParams().
		WithAppName(appName)

	clientBearerToken := httptransport.BearerToken(machineUserToken)
	client := httpUtils.GetApplicationClient(cfg)

	application, err := client.GetApplication(params, clientBearerToken)
	if err != nil {
		return nil, err
	}

	return application.Payload, nil
}

func buildApp(cfg config.Config, machineUserToken string) error {
	bodyParameters := models.PipelineParametersBuild{
		Branch: defaults.App2BranchToBuildFrom,
	}

	params := apiclient.NewTriggerPipelineBuildParams().
		WithAppName(defaults.App2Name).
		WithPipelineParametersBuild(&bodyParameters)

	clientBearerToken := httptransport.BearerToken(machineUserToken)
	client := httpUtils.GetApplicationClient(cfg)

	_, err := client.TriggerPipelineBuild(params, clientBearerToken)
	if err != nil {
		return err
	}

	return nil
}

func setSecret(cfg config.Config, machineUserToken string) error {
	params := environmentclient.NewChangeComponentSecretParams().
		WithAppName(defaults.App2Name).
		WithEnvName(defaults.App2EnvironmentName).
		WithComponentName(defaults.App2Component2Name).
		WithSecretName(defaults.App2SecretName).
		WithComponentSecret(
			&models.SecretParameters{
				SecretValue: stringPtr(defaults.App2SecretValue),
			})

	clientBearerToken := httptransport.BearerToken(machineUserToken)
	client := httpUtils.GetEnvironmentClient(cfg)

	_, err := client.ChangeComponentSecret(params, clientBearerToken)
	if err != nil {
		return fmt.Errorf("failed calling ChangeComponentSecret for application %s: %w", defaults.App2Name, err)
	}
	return nil
}

func (s *step) isTriggerPipelineBuildUnauthorized(err error) bool {
	return err != nil && s.checkErrorResponse(err, 401)
}

func (s *step) isGetApplicationForbidden(err error) bool {
	if errors.Is(err, &apiclient.GetApplicationForbidden{}) {
		return true
	}
	s.logger.Debugf("expected err apiclient.GetApplicationForbidden but got: %v", err)
	return false
}

func (s *step) checkErrorResponse(err error, expectedStatusCode int) bool {
	apiError, ok := err.(*runtime.APIError)
	if ok {
		errorCode := apiError.Code
		s.logger.Debugf("checkErrorResponse err code: %d, expected err code: %d", errorCode, expectedStatusCode)
		if errorCode == expectedStatusCode {
			return true
		}
	} else {
		s.logger.Debugf("checkErrorResponse err is not runtime.APIError")
	}
	return false
}

func stringPtr(str string) *string {
	return &str
}
