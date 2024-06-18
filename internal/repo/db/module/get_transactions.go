package module

import (
	"fmt"
	"tx-bank/internal/model/transaction"
)

func (r *repo) GetTransactions(status, offset, limit int) ([]transaction.TransactionDB, int32, error) {
	var (
		results []transaction.TransactionDB
		query   string
		err     error
	)

	query, params := buildQueryGetTransactions(status, offset, limit)
	rows, err := r.db.Query(query, params...)
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
	params = []interface{}{}
	query = "SELECT COUNT(*) FROM transactions"
	if status > 0 {
		query += " WHERE status=$1;"
		params = append(params, status)
	}
	err = r.db.QueryRow(query, params...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return results, total, err
}

func buildQueryGetTransactions(status, offset, limit int) (string, []interface{}) {
	var (
		query       string = qGetTransactions
		params      []interface{}
		countParams int = 1
	)

	if status > 0 {
		query = query + fmt.Sprintf(" WHERE status = $%d", countParams)
		params = append(params, status)
		countParams++
	}
	if offset > 0 {
		query = query + fmt.Sprintf(" OFFSET $%d", countParams)
		params = append(params, offset)
		countParams++
	}
	if limit > 0 {
		query = query + fmt.Sprintf(" LIMIT $%d", countParams)
		params = append(params, limit)
		countParams++
	}

	return query, params
}
