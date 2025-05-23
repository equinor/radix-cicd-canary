package deployonly

import (
	"context"

	"github.com/equinor/radix-cicd-canary/metrics"
	"github.com/equinor/radix-cicd-canary/scenarios/deployonly/delete"
	"github.com/equinor/radix-cicd-canary/scenarios/deployonly/deploy"
	"github.com/equinor/radix-cicd-canary/scenarios/deployonly/privateimagehub"
	"github.com/equinor/radix-cicd-canary/scenarios/deployonly/register"
	"github.com/equinor/radix-cicd-canary/scenarios/deployonly/teardown"
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
				Name:        "DeleteApplications",
				Description: "Delete applications",
				Test:        delete.Applications,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
		},
		Teardown: []test.Spec{
			{
				Name:        "DeleteDeployOnlyApplication",
				Description: "Delete applications",
				Test:        teardown.TearDown,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
		},
	}
}

func successFunction(ctx context.Context, testName string) {
	metrics.AddTestOne(testName, metrics.Success)
	metrics.AddTestZero(testName, metrics.Errors)
}

func failFunction(ctx context.Context, testName string) {
	metrics.AddTestZero(testName, metrics.Success)
	metrics.AddTestOne(testName, metrics.Errors)
}
