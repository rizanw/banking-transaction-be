package module

import (
	"database/sql"
	"errors"
	"tx-bank/internal/model/transaction"
)

func (r *repo) FindTransaction(transactionID int64) (transaction.TransactionDB, error) {
	var (
		tx  transaction.TransactionDB
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
		return transaction.TransactionDB{}, err
	}

	return tx, nil
}
