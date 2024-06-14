package module

import "tx-bank/internal/model/user"

func (r *repo) InsertUser(in user.UserDB) (int64, error) {
	var (
		userID int64
		err    error
	)

	err = r.db.QueryRow(qInsertUser,
		in.Username,
		in.Password,
		in.Email,
		in.Phone,
		in.CorporateID,
		in.Role,
	).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
