package main

import (
	"log"
	rDB "tx-bank/internal/repo/db"
	psql "tx-bank/internal/repo/db/module"
)

type Repo struct {
	db rDB.DB
}

func newRepo() *Repo {
	db, err := psql.New()
	if err != nil {
		log.Println("!Error init db:", err)
		panic(err)
		return nil
	}

	return &Repo{
		db: db,
	}
}
