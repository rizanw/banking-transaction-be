package user

import "errors"

const (
	RoleNone = iota
	RoleMaker
	RoleApprover
)

func (user UserDB) ValidateRole() error {
	switch user.Role {
	case RoleNone:
		return errors.New("invalid role")
	case RoleApprover:
	case RoleMaker:
	default:
		return errors.New("invalid role")
	}

	return nil
}
