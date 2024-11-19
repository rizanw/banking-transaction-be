package main

import (
	"tx-bank/internal/domain/corporate"
	"tx-bank/internal/domain/otp"
	"tx-bank/internal/domain/transaction"
	"tx-bank/internal/domain/user"
	"tx-bank/internal/infra"
	repocorp "tx-bank/internal/repo/corporate"
	repootp "tx-bank/internal/repo/otp"
	repotransaction "tx-bank/internal/repo/transaction"
	repouser "tx-bank/internal/repo/user"
)

type Repo struct {
	User        user.Repository
	Corporate   corporate.Repository
	Otp         otp.Repository
	Transaction transaction.Repository
}

func newRepo(res *infra.Resources) *Repo {
	return &Repo{
		User:        repouser.New(res.Database),
		Corporate:   repocorp.New(res.Database),
		Otp:         repootp.New(res.Database),
		Transaction: repotransaction.New(res.Database, res.Redis),
	}
}
