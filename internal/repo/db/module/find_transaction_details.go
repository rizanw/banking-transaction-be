package module

import "tx-bank/internal/model/transaction"

func (r *repo) FindTransactionDetails(transactionID int64) ([]transaction.TransactionDetailDB, error) {
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
			return nil, err
		}
		results = append(results, detail)
	}

	return results, nil
}
