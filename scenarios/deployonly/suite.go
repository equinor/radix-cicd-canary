package deployonly

import (
	"github.com/equinor/radix-cicd-canary/metrics"
	"github.com/equinor/radix-cicd-canary/scenarios/deployonly/alias"
	"github.com/equinor/radix-cicd-canary/scenarios/deployonly/delete"
	"github.com/equinor/radix-cicd-canary/scenarios/deployonly/deploy"
	"github.com/equinor/radix-cicd-canary/scenarios/deployonly/privateimagehub"
	"github.com/equinor/radix-cicd-canary/scenarios/deployonly/register"
	"github.com/equinor/radix-cicd-canary/scenarios/test"
)

// TestSuite Get the suite
func TestSuite() test.Suite {
	return test.Suite{
		Name: "Deploy only",
		Setup: []test.Spec{
			{
				Name:        "RegisterDeployOnlyApplication",
				Description: "Register application",
				Test:        register.Application,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
		},
		Tests: []test.Spec{
			{
				Name:        "DeployDeployOnlyApplication",
				Description: "Deploy application",
				Test:        deploy.Application,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
			{
				Name:        "SetDeployOnlyPrivateImageHub",
				Description: "Check private image hub func",
				Test:        privateimagehub.Set,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
			{
				Name:        "DefaultDeployOnlyAliasResponding",
				Description: "Check alias responding",
				Test:        alias.DefaultResponding,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
		},
		Teardown: []test.Spec{
			{
				Name:        "DeleteDeployOnlyApplication",
				Description: "Delete applications",
				Test:        delete.Applications,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
		},
	}
}

func successFunction(testName string) {
	metrics.AddTestSuccess(testName, metrics.Success)
	metrics.AddTestNoError(testName, metrics.Errors)
}

func failFunction(testName string) {
	metrics.AddTestNoSuccess(testName, metrics.Success)
	metrics.AddTestError(testName, metrics.Errors)
}
