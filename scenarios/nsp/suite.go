package nsp

import (
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
			{
				Name:        "MakeInternalDnsLookup",
				Description: "Make DNS lookup toward internal k8s DNS via canary",
				Test:        egresspolicy.LookupInternalDNS,
				SuccessFn:   egresspolicy.InternalDnsSuccess,
				FailFn:      egresspolicy.InternalDnsFail,
			},
			{
				Name:        "MakePublicDnsLookup",
				Description: "Make DNS lookup toward public DNS via canary",
				Test:        egresspolicy.LookupPublicDNS,
				SuccessFn:   egresspolicy.PublicDnsSuccess,
				FailFn:      egresspolicy.PublicDnsFail,
			},
			{
				Name:        "GetJobList",
				Description: "Get list of jobs from job scheduler via canary",
				Test:        egresspolicy.GetJobList,
				SuccessFn:   egresspolicy.GetJobListSuccess,
				FailFn:      egresspolicy.GetJobListFail,
			},
			{
				Name:        "ReachOauthIdp",
				Description: "Reach Oauth IDP from Oauth aux pod",
				Test:        egresspolicy.ReachOauthIdp,
				SuccessFn:   egresspolicy.ReachOauthIdpSuccess,
				FailFn:      egresspolicy.ReachOauthIdpFail,
			},
		},
		Teardown: []test.Spec{},
	}
}
