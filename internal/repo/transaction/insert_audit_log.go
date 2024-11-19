package transaction

import (
	"context"
	domain "tx-bank/internal/domain/transaction"
)

const qInsertAuditLog = `
	INSERT INTO "audit_logs"
		("id", "transaction_id", "user_id", "action", "timestamp")
	VALUES 
		($1,$2,$3,$4,$5)
	RETURNING id;
`

func (r *repo) InsertAuditLog(ctx context.Context, log *domain.AuditLog) error {
	_, err := r.db.QueryxContext(
		ctx,
		qInsertAuditLog,
		log.ID,
		log.TransactionID,
		log.UserID,
		log.Action,
		log.Timestamp,
	)
	if err != nil {
		return err
	}

	return nil
}
