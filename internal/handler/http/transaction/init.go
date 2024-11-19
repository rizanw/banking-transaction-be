package transaction

import (
	ucTransaction "tx-bank/internal/usecase/transaction"
)

type Handler struct {
	ucTransaction ucTransaction.UseCase
}

func New(ucTransaction ucTransaction.UseCase) *Handler {
	return &Handler{
		ucTransaction: ucTransaction,
	}
}
