package user

import (
	"context"

	"github.com/google/uuid"
)

//go:generate
type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	FindUsers(ctx context.Context, username, email *string, corporateID *uuid.UUID) ([]User, error)
}
