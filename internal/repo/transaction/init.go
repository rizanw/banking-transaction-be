package transaction

import (
	dtransaction "tx-bank/internal/domain/transaction"
	"tx-bank/internal/infra/cache"

	"github.com/jmoiron/sqlx"
)

type repo struct {
	db    *sqlx.DB
	cache cache.Cache
}

func New(db *sqlx.DB, cache cache.Cache) dtransaction.Repository {
	return &repo{
		db:    db,
		cache: cache,
	}
}
