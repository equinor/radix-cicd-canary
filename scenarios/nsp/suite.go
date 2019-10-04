package nsp

import (
	"github.com/equinor/radix-cicd-canary/scenarios/nsp/ingress"
	"github.com/equinor/radix-cicd-canary/scenarios/nsp/service"
	"github.com/equinor/radix-cicd-canary/scenarios/test"
)

// TestSuite Get the suite
func TestSuite() test.Suite {
	return test.Suite{
		Name:  "NSP",
		Setup: []test.Spec{},
		Tests: []test.Spec{
			{
				Name:        "ReachIngress",
				Description: "Reach ingress with HTTP GET",
				Test:        ingress.Reach,
				SuccessFn:   ingress.Success,
				FailFn:      ingress.Fail,
			},
			{
				Name:        "ReachServiceDifferentNamespace",
				Description: "Reach service in different namespace",
				Test:        service.Reach,
				SuccessFn:   service.Success,
				FailFn:      service.Fail,
			},
		},
		Teardown: []test.Spec{},
	}
}
