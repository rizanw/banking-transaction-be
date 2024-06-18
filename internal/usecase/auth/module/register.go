package module

import (
	"errors"
	"time"
	"tx-bank/internal/model/auth"
	"tx-bank/internal/model/corporate"
	"tx-bank/internal/model/user"

	"golang.org/x/crypto/bcrypt"
)

func (u usecase) Register(in auth.RegisterRequest) error {
	var (
		now      = time.Now()
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

	// validate otp
	if otpData, err := u.db.FindOTP(in.Email, in.Code); err != nil {
		return err
	} else if otpData.ID == 0 {
		return errors.New("OTP code is invalid")
	} else if now.After(otpData.Expire) {
		return errors.New("OTP code is expired")
	}

	// validate and corporate data
	if corpData, err = u.db.FindCorporate(0, in.CorporateAccountNumber); err != nil {
		return err
	} else if corpData.ID == 0 {
		return errors.New("corporate not found")
	}

	// validate is user exist
	if users, err := u.db.FindUsers(in.Username, in.Email, 0, 0); err != nil {
		return err
	} else if len(users) > 0 {
		return errors.New("user already exists")
	}

	// hash password
	pwd, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// add new user
	userData = user.UserDB{
		Username:    in.Username,
		Password:    string(pwd),
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
