package nsp

import (
	"github.com/equinor/radix-cicd-canary/scenarios/nsp-long/egresspolicy"
	"github.com/equinor/radix-cicd-canary/scenarios/test"
)

// TestSuite Get the suite
func TestSuite() test.Suite {
	return test.Suite{
		Name:  "NSP-Long",
		Setup: []test.Spec{},
		Tests: []test.Spec{
			{
				Name:        "StartAndCheckJobBatch",
				Description: "Start job batch. Check that it was created",
				Test:        egresspolicy.StartAndCheckJobBatch,
				SuccessFn:   egresspolicy.StartAndCheckJobBatchSuccess,
				FailFn:      egresspolicy.StartAndCheckJobBatchFail,
			},
			{
				Name:        "ReachRadixSite",
				Description: "Reach another Radix site from canary",
				Test:        egresspolicy.ReachRadixSite,
				SuccessFn:   egresspolicy.ReachRadixSiteSuccess,
				FailFn:      egresspolicy.ReachRadixSiteFail,
			},
			{
				Name:        "NotReachRadixSite",
				Description: "Not reach another Radix site from canary",
				Test:        egresspolicy.NotReachRadixSite,
				SuccessFn:   egresspolicy.NotReachRadixSiteSuccess,
				FailFn:      egresspolicy.NotReachRadixSiteFail,
			},
			{
				Name:        "NotReachExternalSite",
				Description: "Not reach an external site from canary",
				Test:        egresspolicy.NotReachExternalSite,
				SuccessFn:   egresspolicy.NotReachExternalSiteSuccess,
				FailFn:      egresspolicy.NotReachExternalSiteFail,
			},
		},
		Teardown: []test.Spec{},
	}
}
