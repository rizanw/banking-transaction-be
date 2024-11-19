package user

import "errors"

const (
	RoleNone     Role = iota
	RoleMaker    Role = iota
	RoleApprover Role = iota
)

type Role int

func (r Role) ValidateRole() error {
	switch r {
	case RoleApprover:
	case RoleMaker:
	default:
		return errors.New("invalid role")
	}

	return nil
}

func (r Role) GetName() string {
	switch r {
	case RoleApprover:
		return "Approver"
	case RoleMaker:
		return "Maker"
	default:
		return ""
	}
}
