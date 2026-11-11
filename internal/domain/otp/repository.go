package otp

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Store(ctx context.Context, otp *OTP) error
	Find(ctx context.Context, email, code string) (*OTP, error)
	Deactivate(ctx context.Context, id uuid.UUID) error
}
