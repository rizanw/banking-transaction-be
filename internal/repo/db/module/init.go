package module

import (
	"database/sql"
	"fmt"
	rDB "tx-bank/internal/repo/db"

	_ "github.com/lib/pq"
)

type database struct {
	db *sql.DB
}

func New() (rDB.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "0.0.0.0", 5656, "txbank", "txbank", "txbank")
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &database{
		db: db,
	}, nil
}
