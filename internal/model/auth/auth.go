package auth

import (
	"errors"
	"regexp"
	"unicode"
)

var (
	reAlphaNumeric = regexp.MustCompile("^[a-zA-Z0-9]+$")
	reNumeric      = regexp.MustCompile("^[0-9]+$")
	rePhoneNumber  = regexp.MustCompile(`^\d{1,3}\.\d{6,15}$`)
)

type AuthResponse struct {
	Message string `json:"message"`
}

func ValidatePassword(password string) error {
	count := 0
	isContainNumber, isContainLetter, isContainSpecial := false, false, false
	for _, c := range password {
		count++
		switch {
		case unicode.IsNumber(c):
			isContainNumber = true
		case unicode.IsLetter(c):
			isContainLetter = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			isContainSpecial = true
		}
	}

	if !isContainNumber || !isContainLetter || !isContainSpecial || count < 8 {
		return errors.New("invalid password")
	}

	return nil
}
