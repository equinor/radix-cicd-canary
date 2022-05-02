package deployonly

import (
	metrics2 "github.com/equinor/radix-cicd-canary/metrics"
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
	metrics2.AddTestSuccess(testName)
	metrics2.AddTestNoError(testName)
}

func failFunction(testName string) {
	metrics2.AddTestNoSuccess(testName)
	metrics2.AddTestError(testName)
}
