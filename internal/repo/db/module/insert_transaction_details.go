package module

import (
	"fmt"
	"tx-bank/internal/model/transaction"
)

func (r *repo) InsertTransactionDetails(in []transaction.TransactionDetailDB) error {
	var (
		query  string
		params []interface{}
		err    error
	)

	query, params = buildQueryInsertTransactionDetails(in)
	_, err = r.db.Exec(query, params...)
	if err != nil {
		return err
	}

	return nil
}

func buildQueryInsertTransactionDetails(in []transaction.TransactionDetailDB) (string, []interface{}) {
	var (
		query  string = qInsertTransactionDetails
		params []interface{}
		offset int = 0
	)

	for _, row := range in {
		query += fmt.Sprintf(
			qInsertTransactionDetailsValues,
			offset+1, offset+2, offset+3, offset+4, offset+5, offset+6, offset+7,
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

		offset += 7
	}

	query = query[0 : len(query)-1]
	return query, params
}
