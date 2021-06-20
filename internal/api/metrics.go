package api

import "github.com/prometheus/client_golang/prometheus"

var (
	multiCreateSolutionV1Metrics = newRequestsMetrics("MultiCreateSolutionV1")
	createSolutionV1Metrics      = newRequestsMetrics("CreateSolutionV1")
	updateSolutionV1Metrics      = newRequestsMetrics("UpdateSolutionV1")
	removeSolutionV1Metrics      = newRequestsMetrics("RemoveSolutionV1")

	multiCreateVerdictV1Metrics = newRequestsMetrics("MultiCreateVerdictV1")
	createVerdictV1Metrics      = newRequestsMetrics("CreateVerdictV1")
	updateVerdictV1Metrics      = newRequestsMetrics("UpdateVerdictV1")
	removeVerdictV1Metrics      = newRequestsMetrics("RemoveVerdictV1")
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(multiCreateSolutionV1Metrics.Total)
	prometheus.MustRegister(multiCreateSolutionV1Metrics.Succeeded)
	prometheus.MustRegister(createSolutionV1Metrics.Total)
	prometheus.MustRegister(createSolutionV1Metrics.Succeeded)
	prometheus.MustRegister(updateSolutionV1Metrics.Total)
	prometheus.MustRegister(updateSolutionV1Metrics.Succeeded)
	prometheus.MustRegister(removeSolutionV1Metrics.Total)
	prometheus.MustRegister(removeSolutionV1Metrics.Succeeded)

	prometheus.MustRegister(multiCreateVerdictV1Metrics.Total)
	prometheus.MustRegister(multiCreateVerdictV1Metrics.Succeeded)
	prometheus.MustRegister(createVerdictV1Metrics.Total)
	prometheus.MustRegister(createVerdictV1Metrics.Succeeded)
	prometheus.MustRegister(updateVerdictV1Metrics.Total)
	prometheus.MustRegister(updateVerdictV1Metrics.Succeeded)
	prometheus.MustRegister(removeVerdictV1Metrics.Total)
	prometheus.MustRegister(removeVerdictV1Metrics.Succeeded)
}

type requestsMetrics struct {
	Total     prometheus.Counter
	Succeeded prometheus.Counter
}

func newRequestsMetrics(rpcName string) *requestsMetrics {
	return &requestsMetrics{
		Total: prometheus.NewCounter(prometheus.CounterOpts{
			Name: rpcName + "TotalRequests",
			Help: "Total number of " + rpcName + " requests",
		}),
		Succeeded: prometheus.NewCounter(prometheus.CounterOpts{
			Name: rpcName + "SucceededRequests",
			Help: "Number of succeeded " + rpcName + " requests",
		}),
	}
}
