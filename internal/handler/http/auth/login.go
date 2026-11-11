package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"tx-bank/internal/dto"
	ucAuth "tx-bank/internal/usecase/auth"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		req dto.LoginRequest
		res struct {
			Message string `json:"message"`
		}
		err error
	)
	w.Header().Set("Content-Type", "application/json")

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = req.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Message = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	resp, err := h.ucAuth.Login(ctx, req)
	if errors.Is(err, ucAuth.ErrInvalidCredentials) {
		w.WriteHeader(http.StatusUnauthorized)
		res.Message = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = "failed to login"
		json.NewEncoder(w).Encode(res)
		return
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
