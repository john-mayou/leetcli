package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/john-mayou/leetcli/handler"
	"github.com/stretchr/testify/require"

	_ "github.com/john-mayou/leetcli/internal/testutil"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/health", nil)
	w := httptest.NewRecorder()

	handler := handler.NewTestHandler(nil)
	handler.HealthCheck(w, req)

	var actual map[string]interface{}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &actual))
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, map[string]interface{}{"status": "ok"}, actual)
}
