package module

import (
	"tx-bank/internal/model/transaction"
)

func (r *repo) GetTransactions(offset, limit int) ([]transaction.TransactionDB, int32, error) {
	var (
		results []transaction.TransactionDB
		err     error
	)

	rows, err := r.db.Query(qGetTransactions, offset, limit)
	for rows.Next() {
		var tx transaction.TransactionDB
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
			return nil, 0, err
		}
		results = append(results, tx)
	}

	var total int32
	err = r.db.QueryRow("SELECT COUNT(*) FROM transactions").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return results, total, err
}
