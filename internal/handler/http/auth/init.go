package auth

import (
	ucAuth "tx-bank/internal/usecase/auth"
)

type Handler struct {
	ucAuth ucAuth.UseCase
}

func New(ucAuth ucAuth.UseCase) *Handler {
	return &Handler{
		ucAuth: ucAuth,
	}
}
