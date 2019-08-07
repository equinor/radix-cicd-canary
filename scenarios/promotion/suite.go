package promotion

import "github.com/equinor/radix-cicd-canary/scenarios/test"

// TestSuite Get the suite
// NOTE: This suite is still dependent on happy-path setup
func TestSuite() test.Suite {
	return test.Suite{
		Name: "Promote deployment",
		Tests: []test.Spec{
			{
				Name:        "PromoteDeploymentToOtherEnvironment",
				Description: "Promote deployment to other environment",
				Test:        promoteDeploymentToAnotherEnvironment,
			},
			{
				Name:        "PromoteDeploymentWithinEnvironment",
				Description: "Promote deployment to same environment",
				Test:        promoteDeploymentWithinEnvironment,
			},
		},
	}
}
