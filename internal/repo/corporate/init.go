package corporate

import (
	domain "tx-bank/internal/domain/corporate"

	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) domain.Repository {
	return &repo{
		db: db,
	}
}
