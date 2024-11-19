package transaction

import (
	"context"
	dtransaction "tx-bank/internal/domain/transaction"
)

const qCountTransactionsGroupedStatus = `
	SELECT status, COUNT(*) as count
	FROM transactions
	GROUP BY status;
`

func (r *repo) CountTransactionsGroupByStatus(ctx context.Context) (*dtransaction.StatusCounter, error) {
	rows, err := r.db.QueryxContext(ctx, qCountTransactionsGroupedStatus)
	if err != nil {
		return &dtransaction.StatusCounter{}, err
	}
	defer rows.Close()

	data := make(map[int32]int64, 0)
	for rows.Next() {
		var (
			status int32
			count  int64
		)
		rows.Scan(&status, &count)
		data[status] = count
	}

	return &dtransaction.StatusCounter{
		AwaitingApproval: data[int32(dtransaction.StatusAwaitingApproval)],
		StatusApproved:   data[int32(dtransaction.StatusApproved)],
		StatusRejected:   data[int32(dtransaction.StatusRejected)],
	}, nil
}
