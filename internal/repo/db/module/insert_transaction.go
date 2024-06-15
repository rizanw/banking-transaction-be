package module

import "tx-bank/internal/model/transactions"

func (r *repo) InsertTransaction(in transactions.TransactionDB) (int64, error) {
	var (
		transactionID int64
		err           error
	)

	err = r.db.QueryRow(qInsertTransaction,
		in.RefNum,
		in.AmountTotal,
		in.RecordTotal,
		in.Maker,
		in.TxDate,
		in.Status,
	).Scan(&transactionID)
	if err != nil {
		return 0, err
	}

	return transactionID, nil
}
