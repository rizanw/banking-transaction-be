package corporate

import (
	"context"
	"database/sql"
	"errors"
	domain "tx-bank/internal/domain/corporate"

	"github.com/google/uuid"
)

const qFindCorporate = `
		SELECT 
			"id",
			"name",
			"account_num"
		FROM
		    "corporates"
		WHERE
		    account_num=$1 OR id=$2;
`

func (r *repo) FindCorporate(ctx context.Context, id uuid.UUID, accountNum string) (*domain.Corporate, error) {
	var corp domain.Corporate

	err := r.db.QueryRowContext(ctx, qFindCorporate, accountNum, id).Scan(
		&corp.ID,
		&corp.Name,
		&corp.AccountNum,
	)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return &corp, nil
}
