package list

import (
	apiclient "github.com/equinor/radix-cicd-canary/generated-client/client/platform"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Applications Test that we are able to list applications
func Applications(env env.Env, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := apiclient.NewShowApplicationsParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetPlatformClient(env)

	showAppOk, err := client.ShowApplications(params, clientBearerToken)
	if err == nil {
		logger.Infof("Response length: %v", len(showAppOk.Payload))
		for i, appSummary := range showAppOk.Payload {
			logger.Infof("App %v: %s", i, appSummary.Name)
		}
	}

	return err
}
