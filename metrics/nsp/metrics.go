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
