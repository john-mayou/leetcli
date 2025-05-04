package handler

import (
	"net/http"

	"github.com/john-mayou/leetcli/internal/httpx"
)

func (h *Handler) GetProblems(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"problems":      h.Store.Problems,
		"problems_meta": h.Store.ProblemsMeta,
	})
}
