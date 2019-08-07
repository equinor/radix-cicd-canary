package happypath

import (
	apiclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

func deleteApplications(env env.Env) (bool, error) {
	success, err := deleteApplication(env, app1Name)
	if !success {
		return false, err
	}

	return deleteApplication(env, app2Name)
}

func deleteApplication(env env.Env, appName string) (bool, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := apiclient.NewDeleteApplicationParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(appName)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	_, err := client.DeleteApplication(params, clientBearerToken)
	if err != nil {
		log.Errorf("Error calling DeleteApplication for application %s: %v", appName, err)
	}

	return err == nil, err
}
