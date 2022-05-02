package nsp

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
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
)

// AddingressReachable increases ingressReachable metrics by 1
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
