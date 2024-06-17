package auth

import (
	"errors"
	"tx-bank/internal/model/corporate"
	"tx-bank/internal/model/user"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken   string                `json:"access_token"`
	CorporateInfo corporate.CorporateDB `json:"corporate_info"`
	UserInfo      user.User             `json:"user_info"`
}

func (r LoginRequest) Validate() error {
	// validate username
	if r.Username == "" {
		return errors.New("username is required")
	}
	if !reAlphaNumeric.MatchString(r.Username) {
		return errors.New("username contains invalid characters")
	}

	// validate password
	if r.Password == "" {
		return errors.New("password is required")
	}
	if err := ValidatePassword(r.Password); err != nil {
		return err
	}

	return nil
}
