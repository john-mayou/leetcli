package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	userID, ok := CtxUserID(r.Context())
	if !ok {
		http.Error(w, "unauthorized: missing user id", http.StatusUnauthorized)
		return
	}

	user, err := h.DBClient.FindUserByID(userID)
	if err != nil {
		http.Error(w, "user not found with id: %v", http.StatusNotFound)
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "could not marshal user to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
