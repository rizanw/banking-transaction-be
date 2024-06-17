package module

import (
	hTransaction "tx-bank/internal/handler/http/transaction"
	ucTransaction "tx-bank/internal/usecase/transaction"
)

type handler struct {
	ucTransaction ucTransaction.UseCase
}

func New(ucTransaction ucTransaction.UseCase) hTransaction.Handler {
	return &handler{
		ucTransaction: ucTransaction,
	}
}
