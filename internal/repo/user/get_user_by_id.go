package user

import (
	"context"
	"database/sql"
	"errors"
	domaincorporate "tx-bank/internal/domain/corporate"
	domain "tx-bank/internal/domain/user"

	"github.com/google/uuid"
)

const qGetUserByID = `
	SELECT 
	    "id",
		"username",
		"password",
		"email",
		"phone",
		"corporate_id",
		"role"
	FROM "users" 
	WHERE id = $1;
`

// GetUserByID returns nil, nil when the user doesn't exist.
func (r *repo) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user := domain.User{Corporate: &domaincorporate.Corporate{}}
	err := r.db.QueryRowContext(ctx, qGetUserByID, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Phone,
		&user.Corporate.ID,
		&user.Role,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
