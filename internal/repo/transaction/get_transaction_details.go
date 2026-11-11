package transaction

import (
	"context"
	"database/sql"
	domain "tx-bank/internal/domain/transaction"

	"github.com/google/uuid"
)

const qGetTransactionDetails = `
	SELECT 
		"id", "to_account_num", "to_account_name", "to_account_bank", "amount", "description", "status"
	FROM 
	    "transaction_details"
	WHERE
	    "transaction_id" = $1
	ORDER BY "created_at", "id";
`

// GetTransactionDetails returns the number of details and the details of a transaction.
func (r *repo) GetTransactionDetails(ctx context.Context, trxID uuid.UUID) (int32, []domain.Detail, error) {
	rows, err := r.db.QueryxContext(ctx, qGetTransactionDetails, trxID)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	details := make([]domain.Detail, 0)
	for rows.Next() {
		var (
			detail      domain.Detail
			description sql.NullString
		)
		if err = rows.Scan(
			&detail.ID,
			&detail.ToAccountNumber,
			&detail.ToAccountName,
			&detail.ToAccountBank,
			&detail.Amount,
			&description,
			&detail.Status,
		); err != nil {
			return 0, nil, err
		}
		detail.Description = description.String
		details = append(details, detail)
	}
	if err = rows.Err(); err != nil {
		return 0, nil, err
	}

	return int32(len(details)), details, nil
}
