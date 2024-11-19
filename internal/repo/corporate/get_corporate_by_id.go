package corporate

import (
	"context"
	domain "tx-bank/internal/domain/corporate"

	"github.com/google/uuid"
)

const qGetCorporateByID = `
		SELECT 
			"id",
			"name",
			"account_num"
		FROM
		    "corporates"
		WHERE
		    id=$1;
`

func (r *repo) GetCorporateByID(ctx context.Context, id uuid.UUID) (*domain.Corporate, error) {
	var corp domain.Corporate

	if err := r.db.QueryRowContext(ctx, qGetCorporateByID, id).Scan(
		&corp.ID,
		&corp.Name,
		&corp.AccountNum,
	); err != nil {
		return nil, err
	}

	return &corp, nil
}
