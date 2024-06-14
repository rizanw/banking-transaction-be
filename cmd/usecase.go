package main

import (
	"tx-bank/internal/config"
	"tx-bank/internal/usecase/auth"
	ucAuth "tx-bank/internal/usecase/auth/module"
)

type UseCase struct {
	Auth auth.UseCase
}

func newUseCase(conf *config.Config, repo *Repo) UseCase {
	return UseCase{
		Auth: ucAuth.New(),
	}
}
