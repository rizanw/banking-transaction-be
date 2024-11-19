package user

import (
	"context"
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

func (r *repo) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var user domain.User

	err := r.db.QueryRowContext(ctx, qGetUserByID, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Phone,
		&user.Corporate.ID,
		&user.Role,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
