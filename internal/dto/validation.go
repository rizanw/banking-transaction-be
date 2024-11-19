package dto

import (
	"errors"
	"unicode"
)

func validatePassword(password string) error {
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
