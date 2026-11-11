package transaction

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	dtransaction "tx-bank/internal/domain/transaction"
	"tx-bank/internal/dto"
	"tx-bank/internal/handler/http/httputil"

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
	if req.TransactionID == uuid.Nil {
		httputil.Message(w, http.StatusBadRequest, "invalid transaction_id")
		return
	}
	userID, ok := httputil.UserID(r)
	if !ok {
		httputil.Message(w, http.StatusUnauthorized, "user not login")
		return
	}
	req.RequesterID = userID

	res, err = h.ucTransaction.GetTransaction(ctx, req)
	if errors.Is(err, dtransaction.ErrNotFound) {
		httputil.Message(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		httputil.Message(w, http.StatusInternalServerError, "failed to get transaction")
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
