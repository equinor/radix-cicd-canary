package deployonly

import (
	metrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/happypath"
	"github.com/equinor/radix-cicd-canary/scenarios/deployonly/build"
	"github.com/equinor/radix-cicd-canary/scenarios/deployonly/delete"
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
			{
				Name:        "BuildDeployOnlyApplication",
				Description: "Build application",
				Test:        build.Application,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
		},
		Tests: []test.Spec{},
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
	metrics.AddTestSuccess(testName)
	metrics.AddTestNoError(testName)
}

func failFunction(testName string) {
	metrics.AddTestNoSuccess(testName)
	metrics.AddTestError(testName)
}
