package transaction

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	domain "tx-bank/internal/domain/transaction"

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

func (r *repo) GetTransactionByID(ctx context.Context, id uuid.UUID) (*domain.Transaction, error) {
	var (
		cacheKey    = fmt.Sprintf("transaction:%s", id.String())
		cacheTTL    = 1 * time.Minute
		transaction domain.Transaction
		err         error
	)

	err = r.cache.Get(ctx, cacheKey, &transaction)
	if err == nil || transaction.ID != uuid.Nil {
		return &transaction, nil
	}

	row := r.db.QueryRowxContext(ctx, qGetTransactionByID, id)
	if err = row.Scan(
		&transaction.ID,
		&transaction.RefNum,
		&transaction.AmountTotal,
		&transaction.RecordTotal,
		&transaction.Maker,
		&transaction.TxDate,
		&transaction.Status,
		&transaction.InstructionType,
		&transaction.CreatedAt,
	); errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	r.cache.Set(ctx, cacheKey, transaction, cacheTTL)

	return &transaction, nil
}
