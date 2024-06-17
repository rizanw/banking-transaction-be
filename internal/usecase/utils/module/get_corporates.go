package module

import "tx-bank/internal/model/corporate"

func (u *usecase) GetCorporates() ([]corporate.CorporateDB, error) {
	corps, err := u.db.GetCorporates()
	if err != nil {
		return nil, err
	}

	return corps, nil
}
