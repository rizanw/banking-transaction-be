package module

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"tx-bank/internal/model/transaction"
)

func (h *handler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		ctx = r.Context()
		req transaction.TransactionRequest
		res transaction.TransactionResponse
	)

	w.Header().Set("Content-Type", "application/json")

	req = parseGetTransactions(r)
	res, err = h.ucTransaction.GetTransactions(ctx, req)
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

	statusParam := r.URL.Query().Get("status")
	status, err := strconv.Atoi(statusParam)
	if err != nil || status < 1 {
		status = 0
	}

	corporateIDParam := r.URL.Query().Get("corporate_id")
	corporateID, err := strconv.ParseInt(corporateIDParam, 10, 64)
	if err != nil || corporateID < 1 {
		corporateID = 0
	}

	startDate := time.Time{}
	startDateParam := r.URL.Query().Get("start_date")
	startDateInt, err := strconv.ParseInt(startDateParam, 10, 64)
	if err != nil || startDateInt < 0 {
		startDate = time.Time{}
	}
	if startDateInt > 0 {
		startDate = time.Unix(startDateInt, 0)
	}

	endDate := time.Time{}
	endDateParam := r.URL.Query().Get("end_date")
	endDateInt, err := strconv.ParseInt(endDateParam, 10, 64)
	if err != nil || endDateInt < 0 {
		endDate = time.Time{}
	}
	if endDateInt > 0 {
		endDate = time.Unix(endDateInt, 0)
	}

	return transaction.TransactionRequest{
		Page:    page,
		PerPage: perPage,
		Filter: transaction.TransactionFilter{
			Status:      status,
			CorporateID: corporateID,
			StartDate:   startDate,
			EndDate:     endDate,
		},
	}
}
