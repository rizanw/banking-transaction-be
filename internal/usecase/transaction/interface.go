package transaction

import (
	"context"
	"tx-bank/internal/model/transaction"
)

type UseCase interface {
	DownloadTemplate() [][]string
	UploadTransaction(in transaction.UploadTransactionRequest, csv []transaction.UploadTransactionCSV) (err error)
	GetTransactions(ctx context.Context, in transaction.TransactionRequest) (transaction.TransactionResponse, error)
	GetTransaction(in transaction.TransactionRequest) (transaction.TransactionDetailResponse, error)
	AuditTransaction(ctx context.Context, trxID int64, action string) error
}
