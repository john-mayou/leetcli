package handler

import (
	"net/http"

	"github.com/john-mayou/leetcli/internal/httpx"
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

	httpx.WriteJSON(w, http.StatusOK, user)
}
