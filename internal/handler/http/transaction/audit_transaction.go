package transaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"tx-bank/internal/domain/transaction"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) AuditTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		ctx    = r.Context()
		userID uuid.UUID
		res    struct {
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

	err = h.ucTransaction.AuditTransaction(ctx, userID, trxID, action)
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

func parseAuditTransaction(r *http.Request) (uuid.UUID, string, error) {
	vars := mux.Vars(r)

	trxIDStr := vars["transactionID"]
	trxID, err := uuid.Parse(trxIDStr)
	if err != nil {
		return uuid.Nil, "", errors.New("invalid transaction_id")
	}

	action := r.URL.Query().Get("action")
	switch action {
	case transaction.AuditActionApprove:
	case transaction.AuditActionReject:
	default:
		return uuid.Nil, "", errors.New("valid action is required")
	}

	return trxID, action, nil
}
