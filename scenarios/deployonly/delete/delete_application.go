package delete

import (
	apiclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Applications Tests that we are able to delete applications
func Applications(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})
	return deleteApplication(env, config.App3Name)
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
		logger.Errorf("Error calling DeleteApplication for application %s: %v", appName, err)
	}

	return err == nil, err
}
