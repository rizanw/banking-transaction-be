package module

import (
	rDB "tx-bank/internal/repo/db"
	ucAuth "tx-bank/internal/usecase/auth"
)

type usecase struct {
	db rDB.Repo
}

func New(db rDB.Repo) ucAuth.UseCase {
	return &usecase{
		db: db,
	}
}
