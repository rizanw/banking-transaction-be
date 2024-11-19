package auth

import (
	"encoding/json"
	"net/http"
	"tx-bank/internal/dto"
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
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
