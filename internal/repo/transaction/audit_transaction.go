package transaction

import (
	"context"
	"time"
	domain "tx-bank/internal/domain/transaction"
)

const (
	qAuditUpdateTransaction = `
	UPDATE "transactions"
	SET "status" = $1, "updated_at" = $2
	WHERE "id" = $3 AND "status" = $4;
`
	qAuditUpdateDetails = `
	UPDATE "transaction_details"
	SET "status" = $1, "updated_at" = $2
	WHERE "transaction_id" = $3;
`
	qAuditInsertLog = `
	INSERT INTO "audit_logs"
		("id", "transaction_id", "user_id", "action", "timestamp")
	VALUES
		($1,$2,$3,$4,$5);
`
)

// AuditTransaction moves a transaction from one status to another, updates its
// details, and records the audit log in a single database transaction. The
// status change only applies if the transaction is still in the `from` status,
// so two approvers can't act on the same transaction; otherwise it returns
// domain.ErrAlreadyAudited.
func (r *repo) AuditTransaction(ctx context.Context, log *domain.AuditLog, from, to domain.Status) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() //nolint:errcheck // no-op after commit

	now := time.Now()
	res, err := tx.ExecContext(ctx, qAuditUpdateTransaction, to, now, log.TransactionID, from)
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err != nil {
		return err
	} else if n == 0 {
		return domain.ErrAlreadyAudited
	}

	if _, err = tx.ExecContext(ctx, qAuditUpdateDetails, to, now, log.TransactionID); err != nil {
		return err
	}
	if _, err = tx.ExecContext(ctx, qAuditInsertLog, log.ID, log.TransactionID, log.UserID, log.Action, log.Timestamp); err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}

	_ = r.cache.Delete(ctx, cacheKeyTransaction(log.TransactionID))
	return nil
}
