package happypath

import (
	"github.com/equinor/radix-cicd-canary/metrics"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/adgroup"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/alias"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/build"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/buildsecrets"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/configbranch"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/delete"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/list"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/privateimagehub"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/promote"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/register"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/secret"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/teardown"
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
			{
				Name:        "RegisterApplicationWithMainConfigBranch",
				Description: "Register application with main as config branch",
				Test:        register.ApplicationWithMainConfigBranch,
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
				Test:        privateimagehub.Set,
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
				Name:        "UnauthorizedReaderAccess",
				Description: "Check that reader access to the application is appropriate",
				Test:        unauthorized.ReaderAccess,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			}, // TODO: move further down after debugging is finished
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
			{
				Name:        "ChangeConfigBranch",
				Description: "Checks that radixconfig.yaml is read from correct config branch",
				Test:        configbranch.Change,
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
				Name:        "DeleteApplication",
				Description: "Delete applications",
				Test:        teardown.TearDown,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
		},
	}
}

func successFunction(testName string) {
	metrics.AddTestOne(testName, metrics.Success)
	metrics.AddTestZero(testName, metrics.Errors)
}

func failFunction(testName string) {
	metrics.AddTestZero(testName, metrics.Success)
	metrics.AddTestOne(testName, metrics.Errors)
}
