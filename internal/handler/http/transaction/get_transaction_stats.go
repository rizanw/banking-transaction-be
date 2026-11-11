package transaction

import (
	"encoding/json"
	"net/http"
	"tx-bank/internal/handler/http/httputil"
)

func (h *Handler) GetTransactionStats(w http.ResponseWriter, r *http.Request) {
	userID, ok := httputil.UserID(r)
	if !ok {
		httputil.Message(w, http.StatusUnauthorized, "user not login")
		return
	}

	resp, err := h.ucTransaction.GetTransactionStats(r.Context(), userID)
	if err != nil {
		httputil.Message(w, http.StatusInternalServerError, "failed to get transaction stats")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
