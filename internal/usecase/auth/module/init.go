package module

import (
	"tx-bank/internal/config"
	rDB "tx-bank/internal/repo/db"
	ucAuth "tx-bank/internal/usecase/auth"
)

type usecase struct {
	db       rDB.Repo
	confJWT  config.JWTConfig
	confSMTP config.SMTPConfig
}

func New(db rDB.Repo, confJWT config.JWTConfig, confSMTP config.SMTPConfig) ucAuth.UseCase {
	return &usecase{
		db:       db,
		confJWT:  confJWT,
		confSMTP: confSMTP,
	}
}
