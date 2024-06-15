package module

import "tx-bank/internal/model/transactions"

func (r *repo) GetTransactions(offset, limit int64) ([]transactions.TransactionDB, error) {
	var (
		results []transactions.TransactionDB
		err     error
	)

	rows, err := r.db.Query(qGetTransactions, offset, limit)
	for rows.Next() {
		var tx transactions.TransactionDB
		err = rows.Scan(
			&tx.ID,
			&tx.RefNum,
			&tx.AmountTotal,
			&tx.RecordTotal,
			&tx.Maker,
			&tx.TxDate,
			&tx.Status,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, tx)
	}

	return results, err
}
