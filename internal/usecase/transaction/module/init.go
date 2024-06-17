package module

import (
	rDB "tx-bank/internal/repo/db"
	ucTransaction "tx-bank/internal/usecase/transaction"
)

type usecase struct {
	db rDB.Repo
}

func New(db rDB.Repo) ucTransaction.UseCase {
	return &usecase{
		db: db,
	}
}
