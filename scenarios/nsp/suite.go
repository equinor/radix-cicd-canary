package nsp

import (
	"github.com/equinor/radix-cicd-canary/metrics"
	"github.com/equinor/radix-cicd-canary/scenarios/nsp/egresspolicy"
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
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
			{
				Name:        "ReachServiceDifferentNamespace",
				Description: "Reach service in different namespace",
				Test:        service.Reach,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
			{
				Name:        "MakeInternalDnsLookup",
				Description: "Make DNS lookup toward internal k8s DNS",
				Test:        egresspolicy.LookupInternalDNS,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
			{
				Name:        "MakePublicDnsLookup",
				Description: "Make DNS lookup toward public DNS",
				Test:        egresspolicy.LookupPublicDNS,
				SuccessFn:   successFunction,
				FailFn:      failFunction,
			},
		},
		Teardown: []test.Spec{},
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
