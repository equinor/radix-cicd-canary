package happypath

import (
	apiclient "github.com/equinor/radix-cicd-canary-golang/generated-client/client/application"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary-golang/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

func deleteApplications() (bool, error) {
	success, err := deleteApplication(app1Name)
	if !success {
		return false, err
	}

	return deleteApplication(app2Name)
}

func deleteApplication(appName string) (bool, error) {
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
		log.Errorf("Error calling DeleteApplication for application %s: %v", appName, err)
	}

	return err == nil, err
}
