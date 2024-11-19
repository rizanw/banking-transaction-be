package user

import (
	"context"
	duser "tx-bank/internal/domain/user"

	"github.com/google/uuid"
)

const qFindUsers = `
	SELECT 
	    "id",
		"username",
		"password",
		"email",
		"phone",
		"corporate_id",
		"role"
	FROM "users" 
	WHERE username = $1 OR email = $2 OR corporate_id = $3;
`

func (r *repo) FindUsers(ctx context.Context, username, email *string, corporateID *uuid.UUID) ([]duser.User, error) {
	var users []duser.User

	rows, err := r.db.QueryContext(ctx, qFindUsers, username, email, corporateID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user duser.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Corporate.ID, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
