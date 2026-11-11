package auth

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrInvalidOTP         = errors.New("invalid or expired OTP code")
	ErrTooManyAttempts    = errors.New("too many invalid OTP attempts, request a new code later")
	ErrUserExists         = errors.New("username or email is already registered")
	ErrCorporateNotFound  = errors.New("corporate account not found")
)
