package module

import (
	"time"
	"tx-bank/internal/model/transaction"
)

func (r *repo) UpdateTransaction(in transaction.TransactionDB) error {
	var (
		err error
	)

	if _, err = r.db.Exec(qUpdateTransaction,
		in.Status,
		time.Now(),
		in.ID,
	); err != nil {
		return err
	}

	return nil
}
