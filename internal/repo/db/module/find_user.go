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
	if err = row.Scan(&res); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return user.UserDB{}, err
	}

	return res, nil
}
