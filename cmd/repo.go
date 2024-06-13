package main

import (
	"log"
	"tx-bank/internal/config"
	rDB "tx-bank/internal/repo/db"
	psql "tx-bank/internal/repo/db/module"
)

type Repo struct {
	db rDB.DB
}

func newRepo(conf *config.Config) *Repo {
	db, err := psql.New(conf.Database)
	if err != nil {
		log.Println("!Error init db:", err)
		panic(err)
		return nil
	}

	return &Repo{
		db: db,
	}
}
