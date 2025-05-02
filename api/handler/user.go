package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(UserIDKey()).(string)
	user, err := h.DBClient.FindUserByID(userID)
	if err != nil {
		http.Error(w, "user not found with id: %v", http.StatusNotFound)
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "could not marshal user to JSON", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
