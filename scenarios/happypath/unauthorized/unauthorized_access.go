package unauthorized

import (
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

// Access Checks that we are not able to enter any application we should not
// have access to
func Access(cfg config.Config, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroup()

	params := application.NewGetApplicationParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(defaults.RestrictedApplicationName)

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetApplicationClient(cfg)

	logger.Debugf("check that inpersonated user ha no access to the application %s", defaults.RestrictedApplicationName)
	_, err := client.GetApplication(params, clientBearerToken)
	return givesAccessError(err)
}

func givesAccessError(err error) error {
	switch err.(type) {
	case *application.GetApplicationForbidden:
		return nil
	}
	return err
}
