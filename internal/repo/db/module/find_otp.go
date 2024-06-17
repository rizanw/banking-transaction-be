package module

import (
	"database/sql"
	"errors"
	"tx-bank/internal/model/auth"
)

func (r *repo) FindOTP(code string, userID int64) (auth.OTP, error) {
	var (
		res auth.OTP
		err error
	)

	row := r.db.QueryRow(qFindOTP, userID, code)
	if err = row.Scan(
		&res.ID, &res.UserID, &res.Code, &res.Expire,
	); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return auth.OTP{}, err
	}

	return res, nil
}
