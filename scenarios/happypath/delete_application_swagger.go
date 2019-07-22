package happypath

import (
	apiclient "github.com/equinor/radix-cicd-canary-golang/generated/client/application"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

func deleteApplicationsSwagger() string {
	const testName = "DeleteApplicationSwagger"

	deleteApplicationSwagger(app1Name, testName)
	deleteApplicationSwagger(app2Name, testName)

	return testName
}

func deleteApplicationSwagger(appName, testName string) {
	const basePath = "/api/v1"

	radixAPIURL := utils.GetRadixAPIURL()
	impersonateUser := utils.GetImpersonateUser()
	impersonateGroup := utils.GetImpersonateGroup()
	bearerToken := utils.GetBearerToken()

	params := apiclient.NewDeleteApplicationParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(appName)
	clientBearerToken := httptransport.BearerToken(bearerToken)
	schemes := []string{"https"}

	transport := httptransport.New(radixAPIURL, basePath, schemes)
	client := apiclient.New(transport, strfmt.Default)

	// TODO
	_, err := client.DeleteApplication(params, clientBearerToken)
	if err != nil {
		addTestError(testName)
		log.Errorf("Error calling DeleteApplication for application %s: %v", appName, err)
	}

}
