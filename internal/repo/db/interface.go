package db

import (
	"tx-bank/internal/model/auth"
	"tx-bank/internal/model/corporate"
	"tx-bank/internal/model/transaction"
	"tx-bank/internal/model/user"
)

type Repo interface {
	InsertCorporate(in corporate.CorporateDB) (int64, error)
	FindCorporate(id int64, accountNum string) (corporate.CorporateDB, error)
	GetCorporates() ([]corporate.CorporateDB, error)
	InsertUser(in user.UserDB) (int64, error)
	FindUser(username, email string, id int64) (user.UserDB, error)
	InsertTransaction(in transaction.TransactionDB) (int64, error)
	GetTransactions(offset, limit int) ([]transaction.TransactionDB, int32, error)
	FindTransaction(transactionID int64) (transaction.TransactionDB, error)
	UpdateTransaction(transaction transaction.TransactionDB) error
	InsertTransactionDetails(in []transaction.TransactionDetailDB) error
	FindTransactionDetails(transactionID int64) ([]transaction.TransactionDetailDB, int32, error)
	UpdateTransactionDetailStatus(trxID int64, status int32) error
	InsertAuditLog(in transaction.AuditLogDB) (int64, error)
	InsertOTP(in auth.OTP) error
	FindOTP(code string, userID int64) (auth.OTP, error)
}
