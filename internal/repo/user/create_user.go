package user

import (
	"context"
	"tx-bank/internal/domain/user"
)

const qInsertUser = `
	INSERT INTO "users"
		("username", "password", "email", "phone", "corporate_id", "role")
	VALUES
		($1,$2,$3,$4,$5,$6)
	RETURNING id;
`

func (r *repo) CreateUser(ctx context.Context, user *user.User) error {
	_, err := r.db.QueryxContext(
		ctx,
		qInsertUser,
		user.Username, user.Password, user.Email, user.Phone, user.Corporate.ID, user.Role,
	)
	if err != nil {
		return err
	}
	return nil
}
