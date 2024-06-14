package module

import "tx-bank/internal/model/corporate"

func (r *repo) InsertCorporate(in corporate.CorporateDB) (int64, error) {
	var (
		corporateID int64
		err         error
	)

	err = r.db.QueryRow(qInsertCorporate, in.AccountNum, in.Name).Scan(&corporateID)
	if err != nil {
		return 0, err
	}

	return corporateID, nil
}
