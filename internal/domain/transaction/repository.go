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
	UpdateTransactionStatus(ctx context.Context, id uuid.UUID, status Status) error
	UpdateTransactionDetailStatus(ctx context.Context, trxID uuid.UUID, status Status) error
	InsertAuditLog(ctx context.Context, log *AuditLog) error
	CountTransactionsGroupByStatus(ctx context.Context) (*StatusCounter, error)
}
