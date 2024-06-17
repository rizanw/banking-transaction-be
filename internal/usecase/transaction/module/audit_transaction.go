package module

import (
	"context"
	"errors"
	"time"
	"tx-bank/internal/common/session"
	"tx-bank/internal/model/transaction"
)

func (u *usecase) AuditTransaction(ctx context.Context, trxID int64, action string) error {
	var (
		trx    transaction.TransactionDB
		status int32
		ses    = ctx.Value("session").(session.Session)
	)

	// convert action to status
	switch action {
	case transaction.AuditActionApprove:
		status = transaction.StatusApproved
	case transaction.AuditActionReject:
		status = transaction.StatusRejected
	}

	// get current trx data
	trx, err := u.db.FindTransaction(trxID)
	if err != nil {
		return err
	}

	// validate transaction:StatusAwaitingApproval only
	if trx.Status != transaction.StatusAwaitingApproval {
		return errors.New("transaction is already changed")
	}

	// update status trx
	trx.Status = status
	err = u.db.UpdateTransaction(trx)
	if err != nil {
		return err
	}

	// update detail trx status
	err = u.db.UpdateTransactionDetailStatus(trxID, status)
	if err != nil {
		return err
	}

	// log audit
	_, err = u.db.InsertAuditLog(transaction.AuditLogDB{
		TransactionID: trxID,
		UserID:        ses.UserID,
		Action:        action,
		Timestamp:     time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}
