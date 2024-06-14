package module

import (
	ucAuth "tx-bank/internal/usecase/auth"
)

type usecase struct {
}

func New() ucAuth.UseCase {
	return &usecase{}
}
