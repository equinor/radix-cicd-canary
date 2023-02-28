package list

import (
	"fmt"

	apiclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/platform"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Applications Test that we are able to list applications
func Applications(cfg config.Config, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroup()

	params := apiclient.NewShowApplicationsParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetPlatformClient(cfg)

	showAppOk, err := client.ShowApplications(params, clientBearerToken)
	if err == nil {
		logger.Infof("Response length: %v", len(showAppOk.Payload))
		for i, appSummary := range showAppOk.Payload {
			logger.Infof("App %v: %s", i, appSummary.Name)
		}
	}

	if len(showAppOk.Payload) == 0 {
		return fmt.Errorf("list of applications returned an empty list")
	}
	return err
}
