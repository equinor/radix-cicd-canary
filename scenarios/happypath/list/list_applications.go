package list

import (
	"fmt"

	apiclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/platform"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger zerolog.Logger

// Applications Test that we are able to list applications
func Applications(cfg config.Config, suiteName string) error {
	logger = log.With().Str("suite", suiteName).Logger()

	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := apiclient.NewShowApplicationsParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	client := httpUtils.GetPlatformClient(cfg)
	showAppOk, err := client.ShowApplications(params, nil)
	if err == nil {
		logger.Info().Msgf("Response length: %v", len(showAppOk.Payload))
		for i, appSummary := range showAppOk.Payload {
			logger.Info().Msgf("App %v: %s", i, appSummary.Name)
		}
	}

	if len(showAppOk.Payload) == 0 {
		return fmt.Errorf("list of applications returned an empty list")
	}
	return err
}
