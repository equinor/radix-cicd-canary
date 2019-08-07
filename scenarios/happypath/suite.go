package happypath

import "github.com/equinor/radix-cicd-canary/scenarios/test"

// TestSuite Get the suite
func TestSuite() test.Suite {
	return test.Suite{
		Name: "Happy path",
		Setup: []test.Spec{
			{
				Name:        "RegisterApplication",
				Description: "Register application",
				Test:        registerApplication,
			},
			{
				Name:        "RegisterApplicationWithNoDeployKey",
				Description: "Register application with no deploy key",
				Test:        registerApplicationWithNoDeployKey,
			},
		},
		Tests: []test.Spec{
			{
				Name:        "ListApplications",
				Description: "List applications",
				Test:        listApplications,
			},
			{
				Name:        "BuildApplication",
				Description: "Build application",
				Test:        buildApplication,
			},
			{
				Name:        "SetSecret",
				Description: "Set secret",
				Test:        setSecret,
			},
			{
				Name:        "DefaultAliasResponding",
				Description: "Check alias responding",
				Test:        defaultAliasResponding,
			},
			{
				Name:        "UnauthorizedAccess",
				Description: "Check access to application user should not be able to access",
				Test:        unauthorizedAccess,
			},
		},
		Teardown: []test.Spec{
			{
				Name:        "DeleteApplication",
				Description: "Delete applications",
				Test:        deleteApplications,
			},
		},
	}
}
