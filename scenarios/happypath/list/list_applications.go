package list

import (
	"context"

	apiclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/platform"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Applications Test that we are able to list applications
func Applications(ctx context.Context, cfg config.Config) error {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := apiclient.NewShowApplicationsParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)

	client := httpUtils.GetPlatformClient(cfg)
	showAppOk, err := client.ShowApplications(params, nil)
	if err == nil {
		log.Ctx(ctx).Info().Msgf("Response length: %v", len(showAppOk.Payload))
		for i, appSummary := range showAppOk.Payload {
			log.Ctx(ctx).Info().Msgf("App %v: %s", i, *appSummary.Name)
		}
	}

	if len(showAppOk.Payload) == 0 {
		return errors.Errorf("list of applications returned an empty list")
	}
	return err
}
