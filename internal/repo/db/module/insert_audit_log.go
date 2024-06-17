package module

import "tx-bank/internal/model/transactions"

func (r *repo) InsertAuditLog(in transactions.AuditLogDB) (int64, error) {
	var (
		id  int64
		err error
	)

	err = r.db.QueryRow(qInsertAuditLog,
		in.TransactionID,
		in.UserID,
		in.Action,
		in.Timestamp,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
