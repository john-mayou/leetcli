package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/john-mayou/leetcli/handler"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"
)

func TestTrackMetric(t *testing.T) {
	tests := map[string]struct {
		Body   handler.TrackMetricBody
		Assert func(t *testing.T, h *handler.Handler)
	}{
		"problem_change_total increments": {
			Body: handler.TrackMetricBody{Name: "problem_change_total", Labels: map[string]int{"problem_number": 1}},
			Assert: func(t *testing.T, h *handler.Handler) {
				require.Equal(t, 1.0, testutil.ToFloat64(h.Metrics.ProblemChangeCounter.WithLabelValues("1")))
			},
		},
		"terminal_focus_total increments": {
			Body: handler.TrackMetricBody{Name: "terminal_focus_total", Labels: map[string]int{"problem_number": 1}},
			Assert: func(t *testing.T, h *handler.Handler) {
				require.Equal(t, 1.0, testutil.ToFloat64(h.Metrics.TerminalFocusCounter.WithLabelValues("1")))
			},
		},
	}

	for tcName, tc := range tests {
		t.Run(tcName, func(t *testing.T) {
			jsonBody, err := json.Marshal(tc.Body)
			require.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/metrics/track", bytes.NewReader(jsonBody))
			w := httptest.NewRecorder()

			handler := handler.NewTestHandler(nil)
			handler.TrackMetric(w, req)

			tc.Assert(t, handler)
		})
	}
}
