package auth

import (
	"encoding/json"
	"net/http"
	"net/mail"
)

func (h *Handler) SendOTP(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		req struct {
			Email string `json:"email"`
		}
		resp struct {
			Message string `json:"message"`
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

	err = h.ucAuth.SendOTP(ctx, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Message = "success, please check your mailbox or spam!"
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
