package transaction

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	CreateTransaction(ctx context.Context, trx *Transaction) error
	GetTransactionByID(ctx context.Context, id uuid.UUID) (*Transaction, error)
	GetTransactions(ctx context.Context, in TransactionFilter) (int32, []Transaction, error)
	GetTransactionDetails(ctx context.Context, trxID uuid.UUID) (int32, []Detail, error)
	AuditTransaction(ctx context.Context, log *AuditLog, from, to Status) error
	CountTransactionsGroupByStatus(ctx context.Context, makers []uuid.UUID) (*StatusCounter, error)
}
