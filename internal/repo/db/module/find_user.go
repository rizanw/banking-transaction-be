package module

import (
	"database/sql"
	"errors"
	"tx-bank/internal/model/user"
)

func (r *repo) FindUser(username, email string) (user.UserDB, error) {
	var (
		res user.UserDB
		err error
	)

	row := r.db.QueryRow(qFindUser, username, email)
	if err = row.Scan(
		&res.ID,
		&res.Username,
		&res.Password,
		&res.Email,
		&res.Phone,
		&res.CorporateID,
		&res.Role,
	); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return user.UserDB{}, err
	}

	return res, nil
}
