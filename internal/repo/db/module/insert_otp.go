package module

import (
	"tx-bank/internal/model/auth"
)

func (r *repo) InsertOTP(in auth.OTP) error {
	var (
		err error
	)

	if _, err = r.db.Exec(qInsertOTP,
		in.Email,
		in.Code,
		in.Expire,
	); err != nil {
		return err
	}

	return nil
}
