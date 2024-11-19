package user

import (
	"context"
	"database/sql"
	"errors"
	domaincorporate "tx-bank/internal/domain/corporate"
	domain "tx-bank/internal/domain/user"

	"github.com/google/uuid"
)

const qGetUserByUsername = `
	SELECT 
	    "id",
		"username",
		"password",
		"email",
		"phone",
		"corporate_id",
		"role"
	FROM "users" 
	WHERE username = $1;
`

type userSchema struct {
	ID          uuid.UUID
	Username    string
	Password    string
	Email       string
	Phone       string
	CorporateID uuid.UUID
	Role        int
}

func (r *repo) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user userSchema

	err := r.db.QueryRowContext(ctx, qGetUserByUsername, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Phone,
		&user.CorporateID,
		&user.Role,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
		Corporate: &domaincorporate.Corporate{
			ID: user.CorporateID,
		},
		Role: domain.Role(user.Role),
	}, nil
}
