package user

import (
	"tx-bank/internal/domain/corporate"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Password  string
	Email     string
	Phone     string
	Corporate *corporate.Corporate
	Role      Role
}

func NewUser(username, password, email, phone string, corporate *corporate.Corporate, role Role) *User {
	return &User{
		ID:        uuid.New(),
		Username:  username,
		Password:  password,
		Email:     email,
		Phone:     phone,
		Corporate: corporate,
		Role:      role,
	}
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) IsExist() bool {
	return len(u.ID.String()) > 0
}
