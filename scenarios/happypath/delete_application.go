package happypath

import (
	apiclient "github.com/equinor/radix-cicd-canary-golang/generated-client/client/application"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary-golang/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

func deleteApplications() string {
	const testName = "DeleteApplication"

	deleteApplication(app1Name, testName)
	deleteApplication(app2Name, testName)

	return testName
}

func deleteApplication(appName, testName string) {
	const basePath = "/api/v1"

	log.Infof("Starting DeleteApplication for application %s...", appName)

	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := apiclient.NewDeleteApplicationParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(appName)

	clientBearerToken := httpUtils.GetClientBearerToken()
	client := httpUtils.GetApplicationClient()

	_, err := client.DeleteApplication(params, clientBearerToken)
	if err != nil {
		addTestError(testName)
		log.Errorf("Error calling DeleteApplication for application %s: %v", appName, err)
	} else {
		addTestSuccess(testName)
		log.Infof("Test success for application %s", appName)
	}

}
