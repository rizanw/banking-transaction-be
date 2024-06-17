package module

import (
	"errors"
	"time"
	"tx-bank/internal/common/session"
	"tx-bank/internal/model/auth"
	"tx-bank/internal/model/user"

	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) Login(in auth.LoginRequest) (auth.LoginResponse, error) {
	var (
		res auth.LoginResponse
		err error
		now = time.Now()
	)

	// find existing user
	userData, err := u.db.FindUser(in.Username, "")
	if err != nil {
		return res, err
	}
	if userData.ID == 0 {
		return res, errors.New("user not found")
	}

	// compare hashed password
	if err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(in.Password)); err != nil {
		return res, errors.New("invalid password")
	}

	// generate token
	token, err := u.generateToken(session.Session{
		UserID: userData.ID,
		Email:  userData.Email,
		Expiry: now.Add(24 * time.Hour).Unix(),
	})
	if err != nil {
		return res, err
	}

	// find corp data
	corpData, err := u.db.FindCorporate(userData.CorporateID, "")
	if err != nil {
		return res, err
	}

	res.AccessToken = token
	res.CorporateInfo = corpData
	res.UserInfo = user.User{
		Username: userData.Username,
		Email:    userData.Email,
		Phone:    userData.Phone,
		Role:     userData.Role,
	}
	return res, nil
}
