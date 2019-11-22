package happypath

import (
	metrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/happypath"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/adgroup"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/alias"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/build"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/buildsecrets"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/delete"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/list"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/privateimagehub"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/promote"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/register"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/secret"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/unauthorized"
	"github.com/equinor/radix-cicd-canary/scenarios/test"
)

// TestSuite Get the suite
func TestSuite() test.Suite {
	return test.Suite{
		Name: "Happy path",
		Setup: []test.Spec{
			{
				Name:        "RegisterApplication",
				Description: "Register application",
				Test:        register.Application,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
			{
				Name:        "RegisterApplicationWithNoDeployKey",
				Description: "Register application with no deploy key",
				Test:        register.ApplicationWithNoDeployKey,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
		},
		Tests: []test.Spec{
			{
				Name:        "ListApplications",
				Description: "List applications",
				Test:        list.Applications,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
			{
				Name:        "BuildSecrets",
				Description: "Set build secrets",
				Test:        buildsecrets.Set,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
			{
				Name:        "BuildApplication",
				Description: "Build application",
				Test:        build.Application,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
			{
				Name:        "SetSecret",
				Description: "Set secret",
				Test:        secret.Set,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
			{
				Name:        "SetPrivateImageHub",
				Description: "Check private image hub func",
				Test:        privateimagehub.PrivateImageHub,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
			{
				Name:        "DefaultAliasResponding",
				Description: "Check alias responding",
				Test:        alias.DefaultResponding,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
			{
				Name:        "UnauthorizedAccess",
				Description: "Check access to application user should not be able to access",
				Test:        unauthorized.Access,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
			{
				Name:        "PromoteDeploymentToOtherEnvironment",
				Description: "Promote deployment to other environment",
				Test:        promote.DeploymentToAnotherEnvironment,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
			{
				Name:        "PromoteDeploymentWithinEnvironment",
				Description: "Promote deployment to same environment",
				Test:        promote.DeploymentWithinEnvironment,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
			{
				Name:        "UpdateADGroup",
				Description: "Checks that access can be locked down by updating AD group",
				Test:        adgroup.Update,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
		},
		Teardown: []test.Spec{
			{
				Name:        "DeleteApplication",
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
