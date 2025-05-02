package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/john-mayou/leetcli/handler"
	"github.com/stretchr/testify/require"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/health", nil)
	w := httptest.NewRecorder()

	handler := handler.NewTestHandler(nil)
	handler.HealthCheck(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, `{"status": "ok"}`, w.Body.String())
}
