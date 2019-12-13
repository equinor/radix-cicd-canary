package list

import (
	apiclient "github.com/equinor/radix-cicd-canary/generated-client/client/platform"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

type returnValue struct {
	showApplicationsOk *apiclient.ShowApplicationsOK
	err                error
}

// Applications Test that we are able to list applications
func Applications(env envUtil.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	ok, result := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return areApplicationListed(env)
	})

	if ok && result.(returnValue).err == nil {
		logger.Infof("Response length: %v", len(result.(returnValue).showApplicationsOk.Payload))
		for i, appSummary := range result.(returnValue).showApplicationsOk.Payload {
			logger.Infof("App %v: %s", i, appSummary.Name)
		}

		return true, nil
	}

	return false, result.(returnValue).err
}

// TODO: Should be one call when RA-1128 is done
func areApplicationListed(env envUtil.Env) (bool, interface{}) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := apiclient.NewShowApplicationsParams().
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetPlatformClient(env)

	showAppOk, err := client.ShowApplications(params, clientBearerToken)

	// Expect to see at least canary 1 and canary 2
	if err == nil && len(showAppOk.Payload) >= 2 {
		return true, returnValue{
			showApplicationsOk: showAppOk,
			err:                err,
		}
	}

	return false, nil

}
