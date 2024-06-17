package transaction

import "tx-bank/internal/model/transaction"

type UseCase interface {
	DownloadTemplate() [][]string
	UploadTransaction(in transaction.UploadTransactionRequest, csv []transaction.UploadTransactionCSV) (err error)
	GetTransactions(in transaction.TransactionRequest) (transaction.TransactionResponse, error)
	GetTransaction(in transaction.TransactionRequest) (transaction.TransactionDetailResponse, error)
}
