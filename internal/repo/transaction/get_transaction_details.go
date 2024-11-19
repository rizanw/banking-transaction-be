package transaction

import (
	"context"
	domain "tx-bank/internal/domain/transaction"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

const qGetTransactionDetails = `
	SELECT 
		"id", "transaction_id", "to_account_num", "to_account_name", "to_account_bank", "amount", "description", "status"
	FROM 
	    "transaction_details"
	WHERE
	    "transaction_id" = $1;
`

func (r *repo) GetTransactionDetails(ctx context.Context, trxID uuid.UUID) (int32, []domain.Detail, error) {
	var (
		eg      errgroup.Group
		details []domain.Detail
		total   int32
	)

	eg.Go(func() error {
		rows, err := r.db.QueryxContext(ctx, qGetTransactionDetails, trxID)
		if err != nil {
			return err
		}
		for rows.Next() {
			var detail domain.Detail
			err = rows.Scan(
				&detail.ID,
				&detail.ToAccountNumber,
				&detail.ToAccountName,
				&detail.ToAccountBank,
				&detail.Amount,
				&detail.Description,
				&detail.Status,
			)
			if err != nil {
				return err
			}
			details = append(details, detail)
		}
		return nil
	})

	eg.Go(func() error {
		if err := r.db.QueryRowxContext(
			ctx,
			"SELECT COUNT(*) FROM transaction_details WHERE transaction_id=$1",
			trxID,
		).Scan(&total); err != nil {
			return err
		}
		return nil
	})

	return total, details, nil
}
