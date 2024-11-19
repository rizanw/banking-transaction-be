package module

import (
	"context"
	dcorporate "tx-bank/internal/domain/corporate"
)

func (u *usecase) GetCorporates(ctx context.Context) ([]dcorporate.Corporate, error) {
	corps, err := u.corporateRepo.GetCorporates(ctx)
	if err != nil {
		return nil, err
	}

	return corps, nil
}
