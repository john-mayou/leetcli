package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/john-mayou/leetcli/internal/metric"
)

type TrackMetricBody struct {
	Name   string         `json:"name"`
	Labels map[string]int `json:"labels"`
}

func (h *Handler) TrackMetric(w http.ResponseWriter, r *http.Request) {
	var body TrackMetricBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON body: %v", err.Error()), http.StatusBadRequest)
		return
	}

	switch body.Name {
	case string(metric.ProblemChangeTotal):
		problem_number, ok := body.Labels["problem_number"]
		if !ok {
			http.Error(w, "Missing 'problem_number' label for problem_change_total metric", http.StatusBadRequest)
			return
		}
		h.Metrics.ProblemChangeCounter.WithLabelValues(strconv.Itoa(problem_number)).Inc()
	case string(metric.TerminalFocusTotal):
		problem_number, ok := body.Labels["problem_number"]
		if !ok {
			http.Error(w, "Missing 'problem_number' label for terminal_focus_total metric", http.StatusBadRequest)
			return
		}
		h.Metrics.TerminalFocusCounter.WithLabelValues(strconv.Itoa(problem_number)).Inc()
	default:
		http.Error(w, fmt.Sprintf("Invalid name parameter: %v", body.Name), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
