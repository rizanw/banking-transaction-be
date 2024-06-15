package module

import (
	"database/sql"
	"errors"
	"tx-bank/internal/model/transactions"
)

func (r *repo) FindTransaction(transactionID int64) (transactions.TransactionDB, error) {
	var (
		tx  transactions.TransactionDB
		err error
	)

	row := r.db.QueryRow(qFindTransaction, transactionID)
	if err = row.Scan(
		&tx.ID,
		&tx.RefNum,
		&tx.AmountTotal,
		&tx.RecordTotal,
		&tx.Maker,
		&tx.TxDate,
		&tx.Status,
	); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return transactions.TransactionDB{}, err
	}

	return tx, nil
}
