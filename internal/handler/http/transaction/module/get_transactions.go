package module

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tx-bank/internal/model/transaction"
)

func (h *handler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		req transaction.TransactionRequest
		res transaction.TransactionResponse
	)

	w.Header().Set("Content-Type", "application/json")

	req = parseGetTransactions(r)
	res, err = h.ucTransaction.GetTransactions(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func parseGetTransactions(r *http.Request) transaction.TransactionRequest {
	pageParam := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	perPageParam := r.URL.Query().Get("per_page")
	perPage, err := strconv.Atoi(perPageParam)
	if err != nil || perPage < 1 {
		perPage = 10
	}

	return transaction.TransactionRequest{
		Page:    page,
		PerPage: perPage,
	}
}
