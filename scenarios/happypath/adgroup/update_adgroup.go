package adgroup

import (
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
	pipelineName        = "build"
	adGroupWithNoAccess = "12345678-9012-3456-7890-123456789012"
)

var logger *log.Entry

// Update Tests that updates to ad group locks down an application
func Update(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	ok, _ := test.WaitForCheckFuncOrTimeout(env, hasAccess)
	if !ok {
		return false, nil
	}

	err := patchAdGroup(env, adGroupWithNoAccess)
	if err != nil {
		return false, err
	}

	ok, _ = test.WaitForCheckFuncOrTimeout(env, hasNoAccess)
	if !ok {
		return false, nil
	}

	err = patchAdGroup(env, env.GetImpersonateGroup())
	if err != nil {
		return false, err
	}

	ok, _ = test.WaitForCheckFuncOrTimeout(env, hasAccess)
	if !ok {
		return false, err
	}

	return true, nil
}

func hasNoAccess(env env.Env) (bool, interface{}) {
	return hasProperAccess(env, false), nil
}

func hasAccess(env env.Env) (bool, interface{}) {
	return hasProperAccess(env, true), nil
}

func hasProperAccess(env env.Env, properAccess bool) bool {
	_, err := getApplication(env)
	accessToApplication := !givesAccessError(err)

	err = buildApp(env)
	accessToBuild := !givesAccessError(err)

	err = setSecret(env)
	accessToSecret := !givesAccessError(err)

	hasProperAccess := accessToApplication == properAccess && accessToBuild == properAccess && accessToSecret == properAccess
	if !hasProperAccess {
		logger.Info("Proper access hasn't been granted yet")
	}

	return hasProperAccess
}

func patchAdGroup(env env.Env, adGroup string) error {
	patchRequest := models.ApplicationPatchRequest{
		AdGroups: []string{adGroup},
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

	bodyParameters := models.PipelineParameters{
		PipelineParametersBuild: models.PipelineParametersBuild{
			Branch: config.App2BranchToBuildFrom,
		},
	}

	params := apiclient.NewTriggerPipelineBuildParams().
		WithAppName(config.App2Name).
		WithPipelineParametersBuild(&bodyParameters.PipelineParametersBuild).
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

	params := environmentclient.NewChangeEnvironmentComponentSecretParams().
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

	_, err := client.ChangeEnvironmentComponentSecret(params, clientBearerToken)
	if err != nil {
		logger.Errorf("Error calling ChangeEnvironmentComponentSecret for application %s: %v", config.App2Name, err)
		return err
	}

	return nil
}

func givesAccessError(err error) bool {
	const successStatusCode = 403
	return err != nil && checkErrorResponse(err, successStatusCode)
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
