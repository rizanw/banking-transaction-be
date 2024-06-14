package db

import (
	"tx-bank/internal/model/corporate"
	"tx-bank/internal/model/user"
)

type Repo interface {
	InsertCorporate(in corporate.CorporateDB) (int64, error)
	FindCorporate(accountNum string) (corporate.CorporateDB, error)
	InsertUser(in user.UserDB) (int64, error)
	FindUser(username, email string) (user.UserDB, error)
}
