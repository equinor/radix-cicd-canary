package happypath

import (
	apiclient "github.com/equinor/radix-cicd-canary-golang/generated-client/client/platform"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary-golang/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

func listApplications() string {
	const (
		testName = "ListApplications"
		basePath = "/api/v1"
	)

	log.Infof("Starting ShowApplications...")

	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := apiclient.NewShowApplicationsParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken()
	client := httpUtils.GetPlatformClient()

	showAppOk, err := client.ShowApplications(params, clientBearerToken)
	if err != nil {
		addTestError(testName)
		log.Errorf("Error calling ShowApplications: %v", err)
	} else {
		addTestSuccess(testName)
		log.Infof("Response length: %v", len(showAppOk.Payload))
		for i, appSummary := range showAppOk.Payload {
			log.Infof("App %v: %s", i, appSummary.Name)
		}
		log.Info("Test success")
	}

	return testName
}
