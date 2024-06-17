package module

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tx-bank/internal/model/transaction"

	"github.com/gorilla/mux"
)

func (h *handler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		req transaction.TransactionRequest
		res transaction.TransactionDetailResponse
	)
	w.Header().Set("Content-Type", "application/json")

	req = parseGetTransaction(r)
	res, err = h.ucTransaction.GetTransaction(req)
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

func parseGetTransaction(r *http.Request) transaction.TransactionRequest {
	vars := mux.Vars(r)
	transactionIDstr := vars["transactionID"]
	transactionID, err := strconv.Atoi(transactionIDstr)
	if err != nil || transactionID < 1 {
		transactionID = 0
	}

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
		TransactionID: int64(transactionID),
		Page:          page,
		PerPage:       perPage,
	}
}
