package module

import (
	"context"
	dtransaction "tx-bank/internal/domain/transaction"
)

func (u *usecase) GetTransactionStats(ctx context.Context) (dtransaction.StatusCounter, error) {
	data, err := u.transactionRepo.CountTransactionsGroupByStatus(ctx)
	if err != nil {
		return dtransaction.StatusCounter{}, err
	}

	return *data, nil
}
