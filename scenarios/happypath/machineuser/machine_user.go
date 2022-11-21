package machineuser

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
	httptransport "github.com/go-openapi/runtime/client"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Create Tests that machine user is created properly
func Create(env envUtil.Env, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Enable machine user
	const enabled = true
	err := patchMachineUser(env, enabled)
	if err != nil {
		return err
	}

	machineUserToken, err := test.WaitForCheckFuncWithValueOrTimeout(env, func(env envUtil.Env) (*string, error) {
		return getMachineUserToken(env)
	}, logger)

	if err != nil {
		return err
	}

	ok, _ := test.WaitForCheckFuncWithValueOrTimeout(env, func(env envUtil.Env) (bool, error) {
		return hasAccess(env, *machineUserToken), nil
	}, logger)

	if !ok {
		return errors.New("does not have expected access with machine token")
	}

	// Should only have access to its own application
	hasAccessToOtherApplication := hasAccessToApplication(env, config.App1Name, *machineUserToken)
	if hasAccessToOtherApplication {
		return fmt.Errorf("has not expected access to another application '%s'", config.App1Name)
	}

	// Disable machine user
	err = patchMachineUser(env, !enabled)
	if err != nil {
		return err
	}

	// Token should no longer have access
	ok, _ = test.WaitForCheckFuncWithValueOrTimeout(env, func(env envUtil.Env) (bool, error) {
		return hasNoAccess(env, *machineUserToken), nil
	}, logger)

	if !ok {
		return errors.New("has not expected access with machine token")
	}
	log.Debug("MachineUser was set and un-set properly")
	return nil
}

func getMachineUserToken(env envUtil.Env) (*string, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := apiclient.NewRegenerateMachineUserTokenParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(config.App2Name)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	tokenResponse, err := client.RegenerateMachineUserToken(params, clientBearerToken)
	if err != nil {
		return nil, fmt.Errorf("cannot regenerate machine token: %v", err)
	}

	return tokenResponse.Payload.Token, nil
}

func patchMachineUser(env envUtil.Env, enabled bool) error {
	log.Debugf("Set MachineUser to %v", enabled)
	patchRequest := models.ApplicationRegistrationPatchRequest{
		ApplicationRegistrationPatch: &models.ApplicationRegistrationPatch{
			MachineUser: &enabled,
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
	log.Debugf("MachineUser has been set to %v", enabled)
	return nil
}

func hasNoAccess(env envUtil.Env, machineUserToken string) bool {
	return hasProperAccess(env, machineUserToken, false)
}

func hasAccess(env envUtil.Env, machineUserToken string) bool {
	return hasProperAccess(env, machineUserToken, true)
}

func hasProperAccess(env envUtil.Env, machineUserToken string, properAccess bool) bool {
	accessToApplication := hasAccessToApplication(env, config.App2Name, machineUserToken)

	err := buildApp(env, machineUserToken)
	accessToBuild := !isTriggerPipelineBuildUnauthorized(err)

	err = setSecret(env, machineUserToken)
	accessToSecret := !isSetSecretUnauthorizedError(err)

	hasProperAccess := accessToApplication == properAccess && accessToBuild == properAccess && accessToSecret == properAccess
	if !hasProperAccess {
		logger.Info("Proper access hasn't been granted yet")
	}

	return hasProperAccess
}

func hasAccessToApplication(env envUtil.Env, appName, machineUserToken string) bool {
	_, err := getApplication(env, appName, machineUserToken)
	return !isGetApplicationUnauthorized(err) && !isGetApplicationForbidden(err)
}

func isGetApplicationUnauthorized(err error) bool {
	if _, ok := err.(*apiclient.GetApplicationUnauthorized); ok {
		return true
	}

	return false
}

func isSetSecretUnauthorizedError(err error) bool {
	if _, ok := err.(*environmentclient.ChangeComponentSecretUnauthorized); ok {
		return true
	}

	return false
}

func getApplication(env envUtil.Env, appName, machineUserToken string) (*models.Application, error) {
	params := apiclient.NewGetApplicationParams().
		WithAppName(appName)

	clientBearerToken := httptransport.BearerToken(machineUserToken)
	client := httpUtils.GetApplicationClient(env)

	application, err := client.GetApplication(params, clientBearerToken)
	if err != nil {
		return nil, err
	}

	return application.Payload, nil
}

func buildApp(env envUtil.Env, machineUserToken string) error {
	bodyParameters := models.PipelineParametersBuild{
		Branch: config.App2BranchToBuildFrom,
	}

	params := apiclient.NewTriggerPipelineBuildParams().
		WithAppName(config.App2Name).
		WithPipelineParametersBuild(&bodyParameters)

	clientBearerToken := httptransport.BearerToken(machineUserToken)
	client := httpUtils.GetApplicationClient(env)

	_, err := client.TriggerPipelineBuild(params, clientBearerToken)
	if err != nil {
		return err
	}

	return nil
}

func setSecret(env envUtil.Env, machineUserToken string) error {
	params := environmentclient.NewChangeComponentSecretParams().
		WithAppName(config.App2Name).
		WithEnvName(config.App2EnvironmentName).
		WithComponentName(config.App2Component2Name).
		WithSecretName(config.App2SecretName).
		WithComponentSecret(
			&models.SecretParameters{
				SecretValue: stringPtr(config.App2SecretValue),
			})

	clientBearerToken := httptransport.BearerToken(machineUserToken)
	client := httpUtils.GetEnvironmentClient(env)

	_, err := client.ChangeComponentSecret(params, clientBearerToken)
	if err != nil {
		return fmt.Errorf("error calling ChangeComponentSecret for application %s: %w", config.App2Name, err)
	}
	return nil
}

func isTriggerPipelineBuildUnauthorized(err error) bool {
	return err != nil && checkErrorResponse(err, 401)
}

func isGetApplicationForbidden(err error) bool {
	switch err.(type) {
	case *apiclient.GetApplicationForbidden:
		return true
	}

	return false
}

func checkErrorResponse(err error, expectedStatusCode int) bool {
	apiError, ok := err.(*runtime.APIError)
	if ok {
		errorCode := apiError.Code
		if errorCode == expectedStatusCode {
			return true
		}
	}
	return false
}

func stringPtr(str string) *string {
	return &str
}
