package module

import (
	hAuth "tx-bank/internal/handler/http/auth"
	ucAuth "tx-bank/internal/usecase/auth"
)

type handler struct {
	ucAuth ucAuth.UseCase
}

func New(ucAuth ucAuth.UseCase) hAuth.Handler {
	return &handler{
		ucAuth: ucAuth,
	}
}
