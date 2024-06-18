package module

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"tx-bank/internal/model/transaction"

	"github.com/gorilla/mux"
)

func (h *handler) AuditTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		ctx = r.Context()
		res struct {
			Message string `json:"message"`
		}
	)
	w.Header().Set("Content-Type", "application/json")

	trxID, action, err := parseAuditTransaction(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Message = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	err = h.ucTransaction.AuditTransaction(ctx, trxID, action)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Message = fmt.Sprintf("successfully %sed", action)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
}

func parseAuditTransaction(r *http.Request) (int64, string, error) {
	vars := mux.Vars(r)
	transactionIDstr := vars["transactionID"]
	transactionID, err := strconv.Atoi(transactionIDstr)
	if err != nil || transactionID < 1 {
		return 0, "", errors.New("invalid transaction_id")
	}

	action := r.URL.Query().Get("action")
	switch action {
	case transaction.AuditActionApprove:
	case transaction.AuditActionReject:
	default:
		return 0, "", errors.New("valid action is required")
	}

	return int64(transactionID), action, nil
}
