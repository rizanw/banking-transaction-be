package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"tx-bank/internal/dto"
	ucAuth "tx-bank/internal/usecase/auth"
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
	switch {
	case errors.Is(err, ucAuth.ErrTooManyAttempts):
		http.Error(w, err.Error(), http.StatusTooManyRequests)
		return
	case errors.Is(err, ucAuth.ErrInvalidOTP), errors.Is(err, ucAuth.ErrCorporateNotFound):
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	case errors.Is(err, ucAuth.ErrUserExists):
		http.Error(w, err.Error(), http.StatusConflict)
		return
	case err != nil && err.Error() == "invalid role":
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	case err != nil:
		http.Error(w, "failed to register", http.StatusInternalServerError)
		return
	}

	resp.Message = "success"
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
