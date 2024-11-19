package module

import (
	"context"
	"errors"
	dtransaction "tx-bank/internal/domain/transaction"

	"github.com/google/uuid"
)

func (u *usecase) AuditTransaction(ctx context.Context, userID, trxID uuid.UUID, action string) error {
	var (
		status dtransaction.Status
	)

	// convert action to status
	switch action {
	case dtransaction.AuditActionApprove:
		status = dtransaction.StatusApproved
	case dtransaction.AuditActionReject:
		status = dtransaction.StatusRejected
	}

	// get current trx data
	trx, err := u.transactionRepo.GetTransactionByID(ctx, trxID)
	if err != nil {
		return err
	}
	if len(trx.ID.String()) <= 0 {
		return errors.New("transaction not found")
	}

	// validate transaction:StatusAwaitingApproval only
	if trx.Status != dtransaction.StatusAwaitingApproval {
		return errors.New("transaction is already changed")
	}

	// update status trx
	err = u.transactionRepo.UpdateTransactionStatus(ctx, trx.ID, status)
	if err != nil {
		return err
	}

	// update detail trx status
	err = u.transactionRepo.UpdateTransactionDetailStatus(ctx, trx.ID, status)
	if err != nil {
		return err
	}

	// log audit
	auditLog := dtransaction.NewAuditLog(trx.ID, userID, action)
	err = u.transactionRepo.InsertAuditLog(ctx, auditLog)
	if err != nil {
		return err
	}

	return nil
}
