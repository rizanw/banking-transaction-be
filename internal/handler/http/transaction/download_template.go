package transaction

import (
	"encoding/csv"
	"net/http"
)

func (h *Handler) DownloadTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=transaction_template.csv")
	w.Header().Set("Content-Type", "text/csv")

	writer := csv.NewWriter(w)
	defer writer.Flush()

	dataCsv := h.ucTransaction.DownloadTemplate()

	err := writer.WriteAll(dataCsv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
