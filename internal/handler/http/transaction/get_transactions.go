package transaction

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	dtransaction "tx-bank/internal/domain/transaction"
	"tx-bank/internal/dto"
	"tx-bank/internal/handler/http/httputil"
)

func (h *Handler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		ctx = r.Context()
		req dto.GetTransactionsRequest
		res dto.GetTransactionsResponse
	)

	w.Header().Set("Content-Type", "application/json")

	req = parseGetTransactionsRequest(r)
	userID, ok := httputil.UserID(r)
	if !ok {
		httputil.Message(w, http.StatusUnauthorized, "user not login")
		return
	}
	req.RequesterID = userID

	res, err = h.ucTransaction.GetTransactions(ctx, req)
	if err != nil {
		httputil.Message(w, http.StatusInternalServerError, "failed to get transactions")
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func parseGetTransactionsRequest(r *http.Request) dto.GetTransactionsRequest {
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
	if perPage > 100 {
		perPage = 100
	}

	statusParam := r.URL.Query().Get("status")
	status, err := strconv.Atoi(statusParam)
	if err != nil || status < 1 {
		status = 0
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

	return dto.GetTransactionsRequest{
		Page:    page,
		PerPage: perPage,
		Filter: dto.FilterTransactionsRequest{
			Status:    dtransaction.Status(status),
			StartDate: startDate,
			EndDate:   endDate,
		},
	}
}
