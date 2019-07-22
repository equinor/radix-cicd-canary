package happypath

import (
	apiclient "github.com/equinor/radix-cicd-canary-golang/generated/client/application"
	models "github.com/equinor/radix-cicd-canary-golang/generated/models"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

func buildApplications() string {
	const testName = "BuildApplications"
	buildApplication(app1Name, testName)

	return testName
}

func buildApplication(appName, testName string) {
	const basePath = "/api/v1"

	radixAPIURL := utils.GetRadixAPIURL()
	impersonateUser := utils.GetImpersonateUser()
	impersonateGroup := utils.GetImpersonateGroup()
	bearerToken := utils.GetBearerToken()

	params := apiclient.NewTriggerPipelineParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(appName).
		WithPipelineName("build-deploy").
		WithPipelineParameters(

			&models.PipelineParameters{
				PipelineParametersBuild: models.PipelineParametersBuild{
					Branch: "master",
				},
			},
		)

	clientBearerToken := httptransport.BearerToken(bearerToken)
	schemes := []string{"https"}

	transport := httptransport.New(radixAPIURL, basePath, schemes)
	client := apiclient.New(transport, strfmt.Default)

	_, err := client.TriggerPipeline(params, clientBearerToken)
	if err != nil {
		addTestError(testName)
		log.Errorf("Error calling Build for application %s: %v", appName, err)
	}

}
