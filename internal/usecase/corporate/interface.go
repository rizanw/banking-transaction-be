package corporate

import (
	"context"
	dcorporate "tx-bank/internal/domain/corporate"
)

type UseCase interface {
	GetCorporates(ctx context.Context) ([]dcorporate.Corporate, error)
}
