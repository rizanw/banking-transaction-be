package module

import (
	"context"
	dtransaction "tx-bank/internal/domain/transaction"

	"github.com/google/uuid"
)

// AuditTransaction approves or rejects a transaction awaiting approval. The
// approver must belong to the same corporate as the transaction's maker.
func (u *usecase) AuditTransaction(ctx context.Context, approverID, trxID uuid.UUID, action string) error {
	var status dtransaction.Status
	switch action {
	case dtransaction.AuditActionApprove:
		status = dtransaction.StatusApproved
	case dtransaction.AuditActionReject:
		status = dtransaction.StatusRejected
	default:
		return dtransaction.ErrInvalidAction
	}

	trx, err := u.transactionRepo.GetTransactionByID(ctx, trxID)
	if err != nil {
		return err
	}
	if trx == nil {
		return dtransaction.ErrNotFound
	}

	allowed, err := u.sameCorporate(ctx, approverID, trx.Maker.ID)
	if err != nil {
		return err
	}
	if !allowed {
		return dtransaction.ErrNotFound
	}

	// the repository only applies the change if the transaction is still awaiting
	// approval, so concurrent approvals can't both succeed
	auditLog := dtransaction.NewAuditLog(trx.ID, approverID, action)
	return u.transactionRepo.AuditTransaction(ctx, auditLog, dtransaction.StatusAwaitingApproval, status)
}
