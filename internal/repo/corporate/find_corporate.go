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

// FindCorporate finds a corporate by ID or account number. It returns nil, nil when none matches.
func (r *repo) FindCorporate(ctx context.Context, id uuid.UUID, accountNum string) (*domain.Corporate, error) {
	var corp domain.Corporate

	err := r.db.QueryRowContext(ctx, qFindCorporate, accountNum, id).Scan(
		&corp.ID,
		&corp.Name,
		&corp.AccountNum,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &corp, nil
}
