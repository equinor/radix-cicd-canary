package adgroup

import (
	"fmt"

	apiclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/client/environment"
	"github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/go-openapi/runtime"
	log "github.com/sirupsen/logrus"
)

const (
	adGroupWithNoAccess = "12345678-9012-3456-7890-123456789012"
)

var logger *log.Entry

// Update Tests that updates to AD group locks down an application
func Update(env env.Env, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	err := test.WaitForCheckFuncOrTimeout(env, hasAccess, logger)
	if err != nil {
		return fmt.Errorf("failed to get update details of the suite %s: %w", suiteName, err)
	}

	err = patchAdGroup(env, adGroupWithNoAccess)
	if err != nil {
		return err
	}

	err = test.WaitForCheckFuncOrTimeout(env, hasNoAccess, logger)
	if err != nil {
		return fmt.Errorf("failed to get patchAdGroup update details: %w", err)
	}

	err = patchAdGroup(env, env.GetImpersonateGroup())
	if err != nil {
		return err
	}

	err = test.WaitForCheckFuncOrTimeout(env, hasAccess, logger)
	if err != nil {
		return err
	}

	return nil
}

func hasNoAccess(env env.Env) error {
	return hasProperAccess(env, false)
}

func hasAccess(env env.Env) error {
	return hasProperAccess(env, true)
}

func hasProperAccess(env env.Env, properAccess bool) error {
	_, err := getApplication(env)
	accessToApplication := !isGetApplicationForbidden(err)

	err = buildApp(env)
	accessToBuild := !isTriggerPipelineBuildForbidden(err)

	err = setSecret(env)
	accessToSecret := !isChangeComponentSecretForbidden(err)

	hasProperAccess := accessToApplication == properAccess && accessToBuild == properAccess && accessToSecret == properAccess
	if !hasProperAccess {
		return fmt.Errorf("proper access hasn't been granted yet")
	}
	return nil
}

func patchAdGroup(env env.Env, adGroup string) error {
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

func getApplication(env env.Env) (*models.Application, error) {
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

func buildApp(env env.Env) error {
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

func setSecret(env env.Env) error {
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

func isChangeComponentSecretForbidden(err error) bool {
	switch err.(type) {
	case *environmentclient.ChangeComponentSecretForbidden:
		return true
	}

	return false
}

func isGetApplicationForbidden(err error) bool {
	switch err.(type) {
	case *apiclient.GetApplicationForbidden:
		return true
	}

	return false
}

func isTriggerPipelineBuildForbidden(err error) bool {
	return err != nil && checkErrorResponse(err, 403)
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
