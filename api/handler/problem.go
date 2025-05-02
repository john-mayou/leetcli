package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetProblems(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{
		"problems":      h.Store.Problems,
		"problems_meta": h.Store.ProblemsMeta,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("error marshaling problem data: %v", err), http.StatusInternalServerError)
	}
}
