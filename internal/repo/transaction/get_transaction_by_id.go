package transaction

import (
	"context"
	"database/sql"
	"errors"
	"time"
	domain "tx-bank/internal/domain/transaction"
	"tx-bank/internal/domain/user"

	"github.com/google/uuid"
)

const qGetTransactionByID = `
	SELECT 
		"id", "ref_num", "amount_total", "record_total", "maker", "date", "status", "instruction_type", "created_at"
	FROM 
	    "transactions"
	WHERE
	    "id" = $1;
`

// GetTransactionByID returns nil, nil when the transaction doesn't exist.
func (r *repo) GetTransactionByID(ctx context.Context, id uuid.UUID) (*domain.Transaction, error) {
	var (
		cacheKey    = cacheKeyTransaction(id)
		cacheTTL    = 1 * time.Minute
		transaction domain.Transaction
		makerID     uuid.UUID
	)

	if err := r.cache.Get(ctx, cacheKey, &transaction); err == nil && transaction.ID != uuid.Nil {
		return &transaction, nil
	}

	err := r.db.QueryRowxContext(ctx, qGetTransactionByID, id).Scan(
		&transaction.ID,
		&transaction.RefNum,
		&transaction.AmountTotal,
		&transaction.RecordTotal,
		&makerID,
		&transaction.TxDate,
		&transaction.Status,
		&transaction.InstructionType,
		&transaction.CreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	transaction.Maker = &user.User{ID: makerID}

	_ = r.cache.Set(ctx, cacheKey, transaction, cacheTTL)

	return &transaction, nil
}
