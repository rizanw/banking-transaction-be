package module

import (
	"context"
	dtransaction "tx-bank/internal/domain/transaction"

	"github.com/google/uuid"
)

// GetTransactionStats counts the requester's corporate transactions per status.
func (u *usecase) GetTransactionStats(ctx context.Context, requesterID uuid.UUID) (dtransaction.StatusCounter, error) {
	_, makers, err := u.corporateMakers(ctx, requesterID)
	if err != nil {
		return dtransaction.StatusCounter{}, err
	}

	data, err := u.transactionRepo.CountTransactionsGroupByStatus(ctx, makers)
	if err != nil {
		return dtransaction.StatusCounter{}, err
	}

	return *data, nil
}
