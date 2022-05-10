package nsp

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	Errors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "radix_nsp_test_errors",
			Help: "Test errors",
		},
		[]string{"testName"},
	)
	Success = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "radix_nsp_test_success",
			Help: "Test success",
		},
		[]string{"testName"},
	)
	serviceUnreachable = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_nsp_service_unreachable",
		Help: "Number of times Service was unreachable",
	})
	serviceReachable = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_nsp_service_reachable",
		Help: "Number of times Service was reachable",
	})
	ingressUnreachable = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_nsp_ingress_unreachable",
		Help: "Number of times Ingress was unreachable",
	})
	// IngressReachable is metrics for counting reachable ingress
	ingressReachable = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_nsp_ingress_reachable",
		Help: "Number of times Ingress was reachable",
	})
	internalDnsLookupSucceeds = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_internal_dns_success",
		Help: "Number of times internal DNS lookup succeeded",
	})
	internalDnsLookupFails = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_internal_dns_fails",
		Help: "Number of times internal DNS lookup failed",
	})
	publicDnsLookupSucceeds = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_public_dns_success",
		Help: "Number of times public DNS lookup succeeded",
	})
	publicDnsLookupFails = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_public_dns_fails",
		Help: "Number of times public DNS lookup failed",
	})
	startAndCheckJobBatchSucceeds = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_job_batch_nsp_success",
		Help: "Number of times creating job batch succeeds",
	})
	startAndCheckJobBatchFails = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_job_batch_nsp_fails",
		Help: "Number of times creating job batch failed",
	})
	getJobListSucceeds = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_job_list_nsp_success",
		Help: "Number of times listing jobs succeeds",
	})
	getJobListFails = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_job_list_nsp_fails",
		Help: "Number of times listing jobs failed",
	})
	reachOauthIdpSucceeds = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_oauth_nsp_success",
		Help: "Number of times reaching oauth through nsp succeeds",
	})
	reachOauthIdpFails = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_oauth_nsp_fails",
		Help: "Number of times reaching oauth through nsp failed",
	})
	reachRadixSiteSucceeds = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_reach_radix_nsp_success",
		Help: "Number of times reaching radix through nsp succeeds",
	})
	reachRadixSiteFails = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_reach_radix_nsp_fails",
		Help: "Number of times reaching radix through nsp failed",
	})
	notReachRadixSiteSucceeds = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_not_reach_radix_nsp_success",
		Help: "Number of times not reaching radix through blocking nsp succeeds",
	})
	notReachRadixSiteFails = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_not_reach_radix_nsp_fails",
		Help: "Number of times not reaching radix through blocking nsp failed",
	})
	notReachExternalSiteSucceeds = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_not_reach_external_site_nsp_success",
		Help: "Number of times not reaching external site through blocking nsp succeeds",
	})
	notReachExternalSiteFails = promauto.NewCounter(prometheus.CounterOpts{
		Name: "radix_not_reach_external_site_nsp_fails",
		Help: "Number of times not reaching external site through blocking nsp failed",
	})
)

// AddIngressReachable increases ingressReachable metrics by 1
func AddIngressReachable() {
	ingressReachable.Inc()
}

// AddIngressUnreachable increases ingressUnreachable metrics by 1
func AddIngressUnreachable() {
	ingressUnreachable.Inc()
}

// AddServiceReachable increases serviceReachable metrics by 1
func AddServiceReachable() {
	serviceReachable.Inc()
}

// AddServiceUnreachable increases serviceUnreachable metrics by 1
func AddServiceUnreachable() {
	serviceUnreachable.Inc()
}

// AddInternalDnsIsHealthy increases internalDnsLookupSucceeds metrics by 1
func AddInternalDnsIsHealthy() {
	internalDnsLookupSucceeds.Inc()
}

// AddInternalDnsIsUnhealthy increases internalDnsLookupFails metrics by 1
func AddInternalDnsIsUnhealthy() {
	internalDnsLookupFails.Inc()
}

// AddPublicDnsIsHealthy increases publicDnsLookupSucceeds metrics by 1
func AddPublicDnsIsHealthy() {
	publicDnsLookupSucceeds.Inc()
}

// AddPublicDnsIsUnhealthy increases publicDnsLookupFails metrics by 1
func AddPublicDnsIsUnhealthy() {
	publicDnsLookupFails.Inc()
}

// AddStartAndCheckJobBatchSuccess increases startAndCheckJobBatchSucceeds metrics by 1
func AddStartAndCheckJobBatchSuccess() {
	startAndCheckJobBatchSucceeds.Inc()
}

// AddStartAndCheckJobBatchFail increases startAndCheckJobBatchFails metrics by 1
func AddStartAndCheckJobBatchFail() {
	startAndCheckJobBatchFails.Inc()
}

// AddGetJobListSuccess increases getJobListSucceeds metrics by 1
func AddGetJobListSuccess() {
	getJobListSucceeds.Inc()
}

// AddGetJobListFail increases getJobListFails metrics by 1
func AddGetJobListFail() {
	getJobListFails.Inc()
}

// AddOauthIdpReachable increases reachOauthIdpSucceeds metrics by 1
func AddOauthIdpReachable() {
	reachOauthIdpSucceeds.Inc()
}

// AddOauthIdpUnreachable increases reachOauthIdpFails metrics by 1
func AddOauthIdpUnreachable() {
	reachOauthIdpFails.Inc()
}

// AddRadixSiteReachable increases reachRadixSiteSucceeds metrics by 1
func AddRadixSiteReachable() {
	reachRadixSiteSucceeds.Inc()
}

// AddRadixSiteUnreachable increases reachRadixSiteFails metrics by 1
func AddRadixSiteUnreachable() {
	reachRadixSiteFails.Inc()
}

// AddNotRadixSiteReachable increases notReachRadixSiteSucceeds metrics by 1
func AddNotRadixSiteReachable() {
	notReachRadixSiteSucceeds.Inc()
}

// AddNotRadixSiteUnreachable increases notReachRadixSiteFails metrics by 1
func AddNotRadixSiteUnreachable() {
	notReachRadixSiteFails.Inc()
}

// AddNotExternalSiteReachable increases notReachRadixSiteFails metrics by 1
func AddNotExternalSiteReachable() {
	notReachExternalSiteSucceeds.Inc()
}

// AddNotExternalSiteUnreachable increases notReachRadixSiteFails metrics by 1
func AddNotExternalSiteUnreachable() {
	notReachExternalSiteFails.Inc()
}
