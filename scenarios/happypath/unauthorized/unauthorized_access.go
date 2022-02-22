package unauthorized

import (
	"github.com/equinor/radix-cicd-canary/generated-client/client/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
)

// Access Checks that we are not able to enter any application we should not
// have access to
func Access(env env.Env, suiteName string) (bool, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := application.NewGetApplicationParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup).
		WithAppName(config.RestrictedApplicationName)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	_, err := client.GetApplication(params, clientBearerToken)
	return givesAccessError(err), nil
}

func givesAccessError(err error) bool {
	switch err.(type) {
	case *application.GetApplicationForbidden:
		return true
	}

	return false
}
