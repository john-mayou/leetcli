package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type MetricName string

const (
	ProblemChangeTotal MetricName = "problem_change_total"
	TerminalFocusTotal MetricName = "terminal_focus_total"
)

type MetricsHandler struct {
	ProblemChangeCounter *prometheus.CounterVec
	TerminalFocusCounter *prometheus.CounterVec
}

func NewMetricsHandler() *MetricsHandler {
	return &MetricsHandler{
		ProblemChangeCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: string(ProblemChangeTotal),
			Help: "Total number of problem changes through dropdown",
		}, []string{"problem_number"}),
		TerminalFocusCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: string(TerminalFocusTotal),
			Help: "Total number of terminal focuses with mouse",
		}, []string{"problem_number"}),
	}
}

func NewTestMetricsHandler() *MetricsHandler {
	return &MetricsHandler{
		ProblemChangeCounter: prometheus.NewCounterVec(prometheus.CounterOpts{Name: string(ProblemChangeTotal)}, []string{"problem_number"}),
		TerminalFocusCounter: prometheus.NewCounterVec(prometheus.CounterOpts{Name: string(TerminalFocusTotal)}, []string{"problem_number"}),
	}
}
