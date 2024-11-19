package corporate

import (
	uccorporate "tx-bank/internal/usecase/corporate"
)

type Handler struct {
	ucCorporate uccorporate.UseCase
}

func New(ucCorporate uccorporate.UseCase) *Handler {
	return &Handler{
		ucCorporate: ucCorporate,
	}
}
