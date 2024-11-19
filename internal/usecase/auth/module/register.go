package module

import (
	"context"
	"errors"
	duser "tx-bank/internal/domain/user"
	"tx-bank/internal/dto"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) Register(ctx context.Context, in dto.RegisterRequest) error {

	// validate role
	role := duser.Role(in.Role)
	if err := role.ValidateRole(); err != nil {
		return err
	}

	// validate otp
	otpData, err := u.otpRepo.Find(ctx, in.Email, in.Code)
	if err != nil {
		return err
	}
	if !otpData.IsValid() {
		return errors.New("OTP is expired")
	}

	// validate and get corporate data
	corporateData, err := u.corporateRepo.FindCorporate(ctx, uuid.Nil, in.CorporateAccountNumber)
	if err != nil {
		return err
	}

	// validate is user exist
	users, err := u.userRepo.FindUsers(ctx, &in.Username, &in.Email, nil)
	if err != nil {
		return err
	}
	if len(users) > 0 {
		return errors.New("user already exists")
	}

	// hash password
	pwd, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// add new user
	user := duser.NewUser(
		in.Username,
		string(pwd),
		in.Email,
		in.Phone,
		corporateData,
		role,
	)
	if err = u.userRepo.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}
