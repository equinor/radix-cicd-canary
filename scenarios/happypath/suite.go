package happypath

import (
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/adgroup"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/alias"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/build"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/delete"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath/list"
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
			},
			{
				Name:        "RegisterApplicationWithNoDeployKey",
				Description: "Register application with no deploy key",
				Test:        register.ApplicationWithNoDeployKey,
			},
		},
		Tests: []test.Spec{
			{
				Name:        "ListApplications",
				Description: "List applications",
				Test:        list.Applications,
			},
			{
				Name:        "BuildApplication",
				Description: "Build application",
				Test:        build.Application,
			},
			{
				Name:        "SetSecret",
				Description: "Set secret",
				Test:        secret.Set,
			},
			{
				Name:        "DefaultAliasResponding",
				Description: "Check alias responding",
				Test:        alias.DefaultResponding,
			},
			{
				Name:        "UnauthorizedAccess",
				Description: "Check access to application user should not be able to access",
				Test:        unauthorized.Access,
			},
			{
				Name:        "PromoteDeploymentToOtherEnvironment",
				Description: "Promote deployment to other environment",
				Test:        promote.DeploymentToAnotherEnvironment,
			},
			{
				Name:        "PromoteDeploymentWithinEnvironment",
				Description: "Promote deployment to same environment",
				Test:        promote.DeploymentWithinEnvironment,
			},
			{
				Name:        "UpdateADGroup",
				Description: "Checks that access can be locked down by upodating AD group",
				Test:        adgroup.Update,
			},
		},
		Teardown: []test.Spec{
			{
				Name:        "DeleteApplication",
				Description: "Delete applications",
				Test:        delete.Applications,
			},
		},
	}
}
