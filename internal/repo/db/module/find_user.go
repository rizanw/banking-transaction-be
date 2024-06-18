package module

import (
	"database/sql"
	"errors"
	"tx-bank/internal/model/user"
)

func (r *repo) FindUsers(username, email string, id, corpId int64) ([]user.UserDB, error) {
	var (
		res []user.UserDB
		err error
	)

	rows, err := r.db.Query(qFindUser, username, email, id, corpId)
	for rows.Next() {
		var usr user.UserDB
		if err = rows.Scan(
			&usr.ID,
			&usr.Username,
			&usr.Password,
			&usr.Email,
			&usr.Phone,
			&usr.CorporateID,
			&usr.Role,
		); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return []user.UserDB{}, err
		}
		res = append(res, usr)
	}

	return res, nil
}
