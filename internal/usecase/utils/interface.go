package utils

import "tx-bank/internal/model/corporate"

type UseCase interface {
	GetCorporates() ([]corporate.CorporateDB, error)
}
