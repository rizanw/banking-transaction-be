package module

import "tx-bank/internal/model/transactions"

func (r *repo) FindTransactionDetails(transactionID int64) ([]transactions.TransactionDetailDB, error) {
	var (
		results []transactions.TransactionDetailDB
		err     error
	)

	rows, err := r.db.Query(qFindTransactionDetails, transactionID)
	for rows.Next() {
		var detail transactions.TransactionDetailDB
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
			return nil, err
		}
		results = append(results, detail)
	}

	return results, nil
}
