package handler

import (
	"net/http"

	"github.com/john-mayou/leetcli/internal/httpx"
)

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
