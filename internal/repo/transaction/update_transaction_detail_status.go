package transaction

import (
	"context"
	"time"
	domain "tx-bank/internal/domain/transaction"

	"github.com/google/uuid"
)

const qUpdateTransactionDetailStatus = `
	UPDATE "transaction_details" 
	SET "status" = $1, "updated_at" = $2 
	WHERE "transaction_id" = $3;
`

func (r *repo) UpdateTransactionDetailStatus(ctx context.Context, trxID uuid.UUID, status domain.Status) error {
	if _, err := r.db.ExecContext(ctx, qUpdateTransactionDetailStatus, status, time.Now(), trxID); err != nil {
		return err
	}

	return nil
}
