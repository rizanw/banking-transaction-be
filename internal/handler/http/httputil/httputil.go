// Package httputil holds small helpers shared by the HTTP handlers.
package httputil

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

// UserID returns the logged-in user's ID set by the auth middleware.
func UserID(r *http.Request) (uuid.UUID, bool) {
	raw, ok := r.Context().Value("user_id").(string)
	if !ok {
		return uuid.Nil, false
	}
	id, err := uuid.Parse(raw)
	if err != nil {
		return uuid.Nil, false
	}
	return id, true
}

// Message writes {"message": msg} with the given status code.
func Message(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": msg})
}
