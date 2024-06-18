package auth

import "tx-bank/internal/model/auth"

type UseCase interface {
	SendOTP(email string) error
	Register(in auth.RegisterRequest) error
	Login(in auth.LoginRequest) (auth.LoginResponse, error)
	Logout(sessionKey string) error
}
