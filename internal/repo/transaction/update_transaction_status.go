package transaction

import (
	"context"
	"time"
	domain "tx-bank/internal/domain/transaction"

	"github.com/google/uuid"
)

const qUpdateTransactionStatus = `
	UPDATE "transactions"
	SET "status" = $1, "updated_at" = $2
	WHERE "id" = $3
`

func (r *repo) UpdateTransactionStatus(ctx context.Context, id uuid.UUID, status domain.Status) error {
	if _, err := r.db.ExecContext(ctx, qUpdateTransactionStatus, status, time.Now(), id); err != nil {
		return err
	}

	return nil
}
