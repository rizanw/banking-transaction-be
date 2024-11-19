package corporate

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetCorporates(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	w.Header().Set("Content-Type", "application/json")

	resp, err := h.ucCorporate.GetCorporates(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
