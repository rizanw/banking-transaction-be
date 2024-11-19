package auth

import (
	"context"
	"tx-bank/internal/dto"
)

type UseCase interface {
	SendOTP(ctx context.Context, email string) error
	Register(ctx context.Context, in dto.RegisterRequest) error
	Login(ctx context.Context, in dto.LoginRequest) (dto.LoginResponse, error)
	Logout(sessionKey string) error
}
