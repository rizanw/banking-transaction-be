package module

import (
	"time"
)

func (r *repo) UpdateTransactionDetailStatus(trxID int64, status int32) error {
	var (
		err error
	)

	if _, err = r.db.Exec(qUpdateTransactionDetailStatus,
		status,
		time.Now(),
		trxID,
	); err != nil {
		return err
	}

	return nil
}
