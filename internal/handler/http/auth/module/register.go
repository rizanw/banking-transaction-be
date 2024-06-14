package module

import (
	"encoding/json"
	"net/http"
	"tx-bank/internal/model/auth"
)

func (h handler) Register(w http.ResponseWriter, r *http.Request) {
	var (
		req auth.RegisterRequest
		err error
	)

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.ucAuth.Register(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(auth.AuthResponse{Message: "success"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
