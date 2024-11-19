package module

import (
	dcorporate "tx-bank/internal/domain/corporate"
	uccorporate "tx-bank/internal/usecase/corporate"
)

type usecase struct {
	corporateRepo dcorporate.Repository
}

func New(corporateRepo dcorporate.Repository) uccorporate.UseCase {
	return &usecase{
		corporateRepo: corporateRepo,
	}
}
