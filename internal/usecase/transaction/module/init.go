package module

import (
	dcoporate "tx-bank/internal/domain/corporate"
	dtransaction "tx-bank/internal/domain/transaction"
	duser "tx-bank/internal/domain/user"
	uctransaction "tx-bank/internal/usecase/transaction"
)

type usecase struct {
	transactionRepo dtransaction.Repository
	userRepo        duser.Repository
	corporateRepo   dcoporate.Repository
}

func New(transactionRepo dtransaction.Repository, userRepo duser.Repository, corporateRepo dcoporate.Repository) uctransaction.UseCase {
	return &usecase{
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
		corporateRepo:   corporateRepo,
	}
}
