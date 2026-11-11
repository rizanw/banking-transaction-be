package transaction

import (
	"context"
	dtransaction "tx-bank/internal/domain/transaction"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const qCountTransactionsGroupedStatus = `
	SELECT status, COUNT(*) as count
	FROM transactions
	WHERE maker = ANY($1)
	GROUP BY status;
`

// CountTransactionsGroupByStatus counts the transactions created by the given makers, per status.
func (r *repo) CountTransactionsGroupByStatus(ctx context.Context, makers []uuid.UUID) (*dtransaction.StatusCounter, error) {
	rows, err := r.db.QueryxContext(ctx, qCountTransactionsGroupedStatus, pq.Array(makers))
	if err != nil {
		return &dtransaction.StatusCounter{}, err
	}
	defer rows.Close()

	data := make(map[int32]int64)
	for rows.Next() {
		var (
			status int32
			count  int64
		)
		if err = rows.Scan(&status, &count); err != nil {
			return &dtransaction.StatusCounter{}, err
		}
		data[status] = count
	}
	if err = rows.Err(); err != nil {
		return &dtransaction.StatusCounter{}, err
	}

	return &dtransaction.StatusCounter{
		AwaitingApproval: data[int32(dtransaction.StatusAwaitingApproval)],
		StatusApproved:   data[int32(dtransaction.StatusApproved)],
		StatusRejected:   data[int32(dtransaction.StatusRejected)],
	}, nil
}
