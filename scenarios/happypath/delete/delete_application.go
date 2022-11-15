package delete

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	commonErrors "github.com/equinor/radix-common/utils/errors"
)

// Applications Tests that we are able to delete applications
func Applications(env env.Env, suiteName string) error {
	var errs []error
	for _, appName := range []string{config.App1Name, config.App2Name, config.App4Name} {
		err := application.Delete(env, appName)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return commonErrors.Concat(errs)
}
