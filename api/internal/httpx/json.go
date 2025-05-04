package httpx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, val interface{}) {
	w.Header().Set("Content-Type", "application/json")

	buf, err := json.Marshal(val)
	if err != nil {
		http.Error(w, fmt.Sprintf("error marshaling JSON response: %q", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	_, _ = w.Write(buf)
}
