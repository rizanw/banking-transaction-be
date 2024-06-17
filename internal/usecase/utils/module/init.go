package module

import (
	rDB "tx-bank/internal/repo/db"
	ucUtils "tx-bank/internal/usecase/utils"
)

type usecase struct {
	db rDB.Repo
}

func New(db rDB.Repo) ucUtils.UseCase {
	return &usecase{
		db: db,
	}
}
