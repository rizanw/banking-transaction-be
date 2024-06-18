package module

import "tx-bank/internal/model/transaction"

func (r *repo) FindTransactionDetails(transactionID int64) ([]transaction.TransactionDetailDB, int32, error) {
	var (
		results []transaction.TransactionDetailDB
		err     error
	)

	rows, err := r.db.Query(qFindTransactionDetails, transactionID)
	for rows.Next() {
		var detail transaction.TransactionDetailDB
		err = rows.Scan(
			&detail.ID,
			&detail.TransactionID,
			&detail.ToAccountNumber,
			&detail.ToAccountName,
			&detail.ToAccountBank,
			&detail.Amount,
			&detail.Description,
			&detail.Status,
		)
		if err != nil {
			return nil, 0, err
		}
		results = append(results, detail)
	}

	var total int32
	err = r.db.QueryRow(
		"SELECT COUNT(*) FROM transaction_details WHERE transaction_id=$1",
		transactionID,
	).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return results, total, nil
}
