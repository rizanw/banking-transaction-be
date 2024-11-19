package transaction

import (
	"context"
	dtransaction "tx-bank/internal/domain/transaction"
	"tx-bank/internal/dto"

	"github.com/google/uuid"
)

type UseCase interface {
	DownloadTemplate() [][]string
	UploadTransaction(ctx context.Context, in dto.UploadTransactionRequest, csv []dto.UploadTransactionCSV) (err error)
	GetTransactions(ctx context.Context, in dto.GetTransactionsRequest) (dto.GetTransactionsResponse, error)
	GetTransaction(ctx context.Context, in dto.GetTransactionRequest) (dto.GetTransactionResponse, error)
	GetTransactionStats(ctx context.Context) (dtransaction.StatusCounter, error)
	AuditTransaction(ctx context.Context, userID, trxID uuid.UUID, action string) error
}
