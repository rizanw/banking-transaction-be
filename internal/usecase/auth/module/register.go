package module

import (
	"errors"
	"tx-bank/internal/model/auth"
	"tx-bank/internal/model/corporate"
	"tx-bank/internal/model/user"
)

func (u usecase) Register(in auth.RegisterRequest) error {
	var (
		err      error
		corpData corporate.CorporateDB
		userData user.UserDB = user.UserDB{
			Role: in.Role,
		}
	)

	// validate role
	if err = userData.ValidateRole(); err != nil {
		return err
	}

	// TODO: validate OTP here

	// validate and corporate data
	if corpData, err = u.db.FindCorporate(in.CorporateAccountNumber); err != nil {
		return err
	} else if corpData.ID == 0 {
		return errors.New("corporate not found")
	}

	// validate is user exist
	if userData, err = u.db.FindUser(in.Username, in.Email); err != nil {
		return err
	} else if userData.ID != 0 {
		return errors.New("user already exists")
	}

	// TODO: hash password here

	// add new user
	userData = user.UserDB{
		Username:    in.Username,
		Password:    in.Password,
		Email:       in.Email,
		Phone:       in.Phone,
		CorporateID: corpData.ID,
		Role:        in.Role,
	}
	if _, err = u.db.InsertUser(userData); err != nil {
		return err
	}

	return nil
}
