package dto

import "regexp"

var (
	reAlphaNumeric = regexp.MustCompile("^[a-zA-Z0-9]+$")
	rePhoneNumber  = regexp.MustCompile(`^\d{1,3}\.\d{6,15}$`)
	reNumeric      = regexp.MustCompile("^[0-9]+$")
)
