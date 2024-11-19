package transaction

import "github.com/google/uuid"

type Detail struct {
	ID              uuid.UUID
	ToAccountNumber string
	ToAccountName   string
	ToAccountBank   string
	Amount          float64
	Description     string
	Status          Status
}

func NewTransactionDetail(toAccountNumber, toAccountName, toAccountBank, desc string, amount float64, status Status) *Detail {
	return &Detail{
		ID:              uuid.New(),
		ToAccountNumber: toAccountNumber,
		ToAccountName:   toAccountName,
		ToAccountBank:   toAccountBank,
		Amount:          amount,
		Description:     desc,
		Status:          status,
	}
}
