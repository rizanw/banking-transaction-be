package module

import (
	"database/sql"
	"fmt"
	"tx-bank/internal/config"
	rDB "tx-bank/internal/repo/db"

	_ "github.com/lib/pq"
)

type repo struct {
	db *sql.DB
}

func New(conf config.DBConfig) (rDB.Repo, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.DBName)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &repo{
		db: db,
	}, nil
}
