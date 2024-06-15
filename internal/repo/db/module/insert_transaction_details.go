package module

import (
	"fmt"
	"tx-bank/internal/model/transactions"
)

func (r *repo) InsertTransactionDetails(in []transactions.TransactionDetailDB) error {
	var (
		query  string
		params []interface{}
		err    error
	)

	query, params = buildQueryInsertTransactionDetails(in)
	r.db.QueryRow(query, params...)

	return err
}

func buildQueryInsertTransactionDetails(in []transactions.TransactionDetailDB) (string, []interface{}) {
	var (
		query  string = qInsertTransactionDetails
		params []interface{}
		offset int = 0
	)

	for _, row := range in {
		query += fmt.Sprintf(
			qInsertTransactionDetailsValues,
			offset+1, offset+2, offset+3, offset+4, offset+5, offset+6,
		)

		params = append(params,
			row.TransactionID,
			row.ToAccountNumber,
			row.ToAccountName,
			row.ToAccountBank,
			row.Amount,
			row.Description,
			row.Status,
		)

		offset += 6
	}

	query = query[0 : len(query)-1]
	return query, params
}
