package main

import (
	"tx-bank/internal/config"
	"tx-bank/internal/usecase/auth"
	ucAuth "tx-bank/internal/usecase/auth/module"
	"tx-bank/internal/usecase/transaction"
	ucTransaction "tx-bank/internal/usecase/transaction/module"
)

type UseCase struct {
	Auth        auth.UseCase
	Transaction transaction.UseCase
}

func newUseCase(conf *config.Config, repo *Repo) UseCase {
	return UseCase{
		Auth:        ucAuth.New(repo.db, conf.JWT),
		Transaction: ucTransaction.New(repo.db),
	}
}
