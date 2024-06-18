package module

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"tx-bank/internal/model/auth"
)

func (h *handler) SendOTP(w http.ResponseWriter, r *http.Request) {
	var (
		req struct {
			Email string `json:"email"`
		}
		err error
	)
	w.Header().Set("Content-Type", "application/json")

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err = mail.ParseAddress(req.Email); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.ucAuth.SendOTP(req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(auth.AuthResponse{Message: "success, please check your mailbox or spam!"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
