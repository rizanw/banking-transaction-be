package module

import (
	"tx-bank/internal/model/corporate"
)

func (r *repo) GetCorporates() ([]corporate.CorporateDB, error) {
	var (
		res []corporate.CorporateDB
		err error
	)

	rows, err := r.db.Query(qGetCorporates)
	if err != nil {
	}
	for rows.Next() {
		var corp corporate.CorporateDB
		err = rows.Scan(
			&corp.ID,
			&corp.Name,
			&corp.AccountNum,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, corp)
	}

	return res, nil
}
