package auth

import (
	"encoding/json"
	"net/http"
	"tx-bank/internal/dto"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		req  dto.RegisterRequest
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

	if err = req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.ucAuth.Register(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Message = "success"
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
