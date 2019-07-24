package happypath

import (
	apiclient "github.com/equinor/radix-cicd-canary-golang/generated-client/client/application"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary-golang/scenarios/utils/http"
	"github.com/go-openapi/runtime"
)

func unauthorizedAccess() (bool, error) {
	const (
		testName          = "UnauthorizedAccess"
		basePath          = "/api/v1"
		successStatusCode = 403
	)

	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := apiclient.NewGetApplicationParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(restrictedApplicationName)

	clientBearerToken := httpUtils.GetClientBearerToken()
	client := httpUtils.GetApplicationClient()

	_, err := client.GetApplication(params, clientBearerToken)
	return err != nil && checkErrorResponse(err, successStatusCode), err
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
