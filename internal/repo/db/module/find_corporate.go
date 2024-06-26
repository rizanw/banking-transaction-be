package module

import (
	"database/sql"
	"errors"
	"tx-bank/internal/model/corporate"
)

func (r *repo) FindCorporate(id int64, accountNum string) (corporate.CorporateDB, error) {
	var (
		res corporate.CorporateDB
		err error
	)

	row := r.db.QueryRow(qFindCorporate, accountNum, id)
	if err = row.Scan(&res.ID, &res.Name, &res.AccountNum); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return corporate.CorporateDB{}, err
	}

	return res, nil
}
