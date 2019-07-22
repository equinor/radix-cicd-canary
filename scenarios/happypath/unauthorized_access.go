package happypath

import (
	apiclient "github.com/equinor/radix-cicd-canary-golang/generated-client/client/application"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

func unauthorizedAccess() string {
	const (
		testName          = "UnauthorizedAccess"
		basePath          = "/api/v1"
		successStatusCode = 403
	)

	log.Infof("Starting GetApplication...")

	radixAPIURL := utils.GetRadixAPIURL()
	impersonateUser := utils.GetImpersonateUser()
	impersonateGroup := utils.GetImpersonateGroup()
	bearerToken := utils.GetBearerToken()

	params := apiclient.NewGetApplicationParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(restrictedApplicationName)
	clientBearerToken := httptransport.BearerToken(bearerToken)
	schemes := []string{"https"}

	transport := httptransport.New(radixAPIURL, basePath, schemes)
	client := apiclient.New(transport, strfmt.Default)

	_, err := client.GetApplication(params, clientBearerToken)
	if err != nil {
		if checkErrorResponse(err, successStatusCode) {
			addTestSuccess(testName)
			log.Info("Test success")
		} else {
			addTestError(testName)
			log.Errorf("Error test %s returned not 403 status code", testName)
		}
	} else {
		addTestError(testName)
		log.Errorf("Error test %s should not return 200 status code", testName)
	}

	return testName
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
