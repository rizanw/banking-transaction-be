package transaction

import "errors"

var (
	ErrNotFound       = errors.New("transaction not found")
	ErrAlreadyAudited = errors.New("transaction has already been approved or rejected")
	ErrInvalidAction  = errors.New("action must be approve or reject")
)
