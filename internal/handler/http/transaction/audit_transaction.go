package transaction

import (
	"errors"
	"fmt"
	"net/http"
	"tx-bank/internal/domain/transaction"
	"tx-bank/internal/handler/http/httputil"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) AuditTransaction(w http.ResponseWriter, r *http.Request) {
	userID, ok := httputil.UserID(r)
	if !ok {
		httputil.Message(w, http.StatusUnauthorized, "user not login")
		return
	}

	trxID, action, err := parseAuditTransaction(r)
	if err != nil {
		httputil.Message(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.ucTransaction.AuditTransaction(r.Context(), userID, trxID, action)
	switch {
	case errors.Is(err, transaction.ErrNotFound):
		httputil.Message(w, http.StatusNotFound, err.Error())
	case errors.Is(err, transaction.ErrAlreadyAudited):
		httputil.Message(w, http.StatusConflict, err.Error())
	case errors.Is(err, transaction.ErrInvalidAction):
		httputil.Message(w, http.StatusBadRequest, err.Error())
	case err != nil:
		httputil.Message(w, http.StatusInternalServerError, "failed to audit transaction")
	default:
		past := map[string]string{transaction.AuditActionApprove: "approved", transaction.AuditActionReject: "rejected"}
		httputil.Message(w, http.StatusOK, fmt.Sprintf("successfully %s", past[action]))
	}
}

func parseAuditTransaction(r *http.Request) (uuid.UUID, string, error) {
	vars := mux.Vars(r)
	trxID, err := uuid.Parse(vars["transactionID"])
	if err != nil {
		return uuid.Nil, "", errors.New("invalid transaction_id")
	}

	action := r.URL.Query().Get("action")
	switch action {
	case transaction.AuditActionApprove, transaction.AuditActionReject:
	default:
		return uuid.Nil, "", transaction.ErrInvalidAction
	}

	return trxID, action, nil
}
