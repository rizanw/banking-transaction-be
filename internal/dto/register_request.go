package dto

import (
	"errors"
	"net/mail"
)

type RegisterRequest struct {
	Username               string `json:"username"`
	Password               string `json:"password"`
	Email                  string `json:"email"`
	Phone                  string `json:"phone"`
	Role                   int32  `json:"role"`
	Code                   string `json:"code"`
	CorporateAccountNumber string `json:"corporate_account_number"`
}

func (r *RegisterRequest) Validate() error {
	// validate username
	if r.Username == "" {
		return errors.New("username is required")
	}
	if len(r.Username) < 6 {
		return errors.New("username must be at least 6 characters")
	}
	if !reAlphaNumeric.MatchString(r.Username) {
		return errors.New("username contains invalid characters")
	}

	// validate email
	if r.Email == "" {
		return errors.New("email is required")
	}
	if _, err := mail.ParseAddress(r.Email); err != nil {
		return errors.New("invalid email")
	}

	// validate password
	if r.Password == "" {
		return errors.New("password is required")
	}
	if err := validatePassword(r.Password); err != nil {
		return err
	}

	// validate phone number
	if r.Phone == "" {
		return errors.New("phone is required")
	}
	if !rePhoneNumber.MatchString(r.Phone) {
		return errors.New("phone contains invalid characters")
	}

	// validate corp account number
	if r.CorporateAccountNumber == "" {
		return errors.New("corporate account number is required")
	}
	if !reNumeric.MatchString(r.CorporateAccountNumber) {
		return errors.New("corporate account number contains invalid characters")
	}

	// validate otp code
	if len(r.Code) != 6 {
		return errors.New("otp code is required")
	}
	if !reNumeric.MatchString(r.Code) {
		return errors.New("otp code contains invalid characters")
	}

	return nil
}
