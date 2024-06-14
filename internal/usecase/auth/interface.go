package auth

import "tx-bank/internal/model/auth"

type UseCase interface {
	Register(in auth.RegisterRequest) error
}
