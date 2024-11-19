package main

import (
	"tx-bank/internal/infra"
	"tx-bank/internal/usecase/auth"
	ucauth "tx-bank/internal/usecase/auth/module"
	"tx-bank/internal/usecase/corporate"
	uccorporate "tx-bank/internal/usecase/corporate/module"
	"tx-bank/internal/usecase/transaction"
	uctransaction "tx-bank/internal/usecase/transaction/module"
)

type UseCase struct {
	Auth        auth.UseCase
	Transaction transaction.UseCase
	Corporate   corporate.UseCase
}

func newUseCase(repo *Repo, res *infra.Resources) UseCase {
	return UseCase{
		Auth:        ucauth.New(res.SessionManager, repo.Otp, repo.User, repo.Corporate, res.NotificationService),
		Transaction: uctransaction.New(repo.Transaction, repo.User, repo.Corporate),
		Corporate:   uccorporate.New(repo.Corporate),
	}
}
