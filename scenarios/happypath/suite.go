package happypath

import "github.com/equinor/radix-cicd-canary-golang/scenarios/test"

// TestSuite Get the suite
func TestSuite() test.Suite {
	return test.Suite{
		Name: "Happy path",
		Tests: []test.Spec{
			{
				Name:        "ListApplications",
				Description: "List applications",
				Test:        listApplications,
			},
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
