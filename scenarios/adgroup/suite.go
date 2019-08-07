package adgroup

import "github.com/equinor/radix-cicd-canary/scenarios/test"

// TestSuite Get the suite
// NOTE: This suite is still dependent on happy-path setup
func TestSuite() test.Suite {
	return test.Suite{
		Name: "Update Ad Group",
		Tests: []test.Spec{
			{
				Name:        "UpdateAdGroup",
				Description: "Update ad group",
				Test:        updateAdGroup,
			},
		},
	}
}
