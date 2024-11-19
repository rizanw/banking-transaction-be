package otp

import (
	"context"
)

type Repository interface {
	Store(ctx context.Context, otp *OTP) error
	Find(ctx context.Context, email, code string) (*OTP, error)
}
