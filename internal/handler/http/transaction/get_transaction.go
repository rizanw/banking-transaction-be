package transaction

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tx-bank/internal/dto"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		err error
		req dto.GetTransactionRequest
		res dto.GetTransactionResponse
	)
	w.Header().Set("Content-Type", "application/json")

	req = parseGetTransaction(r)
	res, err = h.ucTransaction.GetTransaction(ctx, req)
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

func parseGetTransaction(r *http.Request) dto.GetTransactionRequest {
	vars := mux.Vars(r)
	transactionIDstr := vars["transactionID"]
	transactionID, err := uuid.Parse(transactionIDstr)
	if err != nil {
		transactionID = uuid.Nil
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

	return dto.GetTransactionRequest{
		TransactionID: transactionID,
		Page:          page,
		PerPage:       perPage,
	}
}
