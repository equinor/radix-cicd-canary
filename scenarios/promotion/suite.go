package promotion

import "github.com/equinor/radix-cicd-canary/scenarios/test"

// TestSuite Get the suite
func TestSuite() test.Suite {
	return test.Suite{
		Name: "Promote deployment",
		Tests: []test.Spec{
			{
				Name:        "PromoteDeployment",
				Description: "Promote deployment",
				Test:        promoteDeployment,
			},
		},
	}
}
