package delete

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
)

// Applications Tests that we are able to delete applications
func Applications(env env.Env, suiteName string) (bool, error) {
	return application.Delete(env, config.App3Name)
}
