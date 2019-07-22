package happypath

import (
	apiclient "github.com/equinor/radix-cicd-canary-golang/generated/client/platform"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

func listApplicationsSwagger() string {
	const (
		testName = "ListApplicationsSwagger"
		basePath = "/api/v1"
	)

	log.Infof("Starting ShowApplications...")

	radixAPIURL := utils.GetRadixAPIURL()
	impersonateUser := utils.GetImpersonateUser()
	impersonateGroup := utils.GetImpersonateGroup()
	bearerToken := utils.GetBearerToken()

	params := apiclient.NewShowApplicationsParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httptransport.BearerToken(bearerToken)
	schemes := []string{"https"}

	transport := httptransport.New(radixAPIURL, basePath, schemes)
	client := apiclient.New(transport, strfmt.Default)

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
