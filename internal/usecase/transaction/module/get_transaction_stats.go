package module

import "tx-bank/internal/model/transaction"

func (u *usecase) GetTransactionStats() (transaction.TransactionStats, error) {
	var (
		transactionStats transaction.TransactionStats
	)

	data, err := u.db.CountTransactionsGroupedStatus()
	if err != nil {
		return transaction.TransactionStats{}, err
	}

	transactionStats.AwaitingApproval = data[transaction.StatusAwaitingApproval]
	transactionStats.Approved = data[transaction.StatusApproved]
	transactionStats.Rejected = data[transaction.StatusRejected]
	return transactionStats, nil
}
