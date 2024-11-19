package corporate

import (
	"context"
	domain "tx-bank/internal/domain/corporate"
)

const qGetCorporates = `
		SELECT 
			"id",
			"name",
			"account_num"
		FROM
		    "corporates";
`

func (r *repo) GetCorporates(ctx context.Context) ([]domain.Corporate, error) {
	var corporates = make([]domain.Corporate, 0)

	rows, err := r.db.QueryContext(ctx, qGetCorporates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var corporate domain.Corporate
		err = rows.Scan(
			&corporate.ID,
			&corporate.Name,
			&corporate.AccountNum,
		)
		if err != nil {
			return nil, err
		}
		corporates = append(corporates, corporate)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return corporates, nil
}
