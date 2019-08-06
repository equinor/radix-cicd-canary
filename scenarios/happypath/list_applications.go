package happypath

import (
	apiclient "github.com/equinor/radix-cicd-canary/generated-client/client/platform"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

func listApplications() (bool, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := apiclient.NewShowApplicationsParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken()
	client := httpUtils.GetPlatformClient()

	showAppOk, err := client.ShowApplications(params, clientBearerToken)
	if err == nil {
		log.Infof("Response length: %v", len(showAppOk.Payload))
		for i, appSummary := range showAppOk.Payload {
			log.Infof("App %v: %s", i, appSummary.Name)
		}
	}

	return err == nil, err
}
