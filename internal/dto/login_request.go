package dto

import "errors"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate() error {
	// validate username
	if r.Username == "" {
		return errors.New("username cannot be empty")
	}
	if len(r.Username) < 6 {
		return errors.New("username must be at least 6 characters")
	}
	if !reAlphaNumeric.MatchString(r.Username) {
		return errors.New("username contains invalid characters")
	}

	// validate password
	if r.Password == "" {
		return errors.New("password cannot be empty")
	}
	passwordValidation := validatePassword(r.Password)
	if passwordValidation != nil {
		return passwordValidation
	}

	return nil
}
