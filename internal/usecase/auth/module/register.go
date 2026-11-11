package module

import (
	"context"
	"fmt"
	dotp "tx-bank/internal/domain/otp"
	duser "tx-bank/internal/domain/user"
	"tx-bank/internal/dto"
	ucAuth "tx-bank/internal/usecase/auth"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) Register(ctx context.Context, in dto.RegisterRequest) error {
	// validate role
	role := duser.Role(in.Role)
	if err := role.ValidateRole(); err != nil {
		return err
	}

	// limit wrong guesses per email, so a 6-digit code can't be brute-forced
	attemptsKey := fmt.Sprintf("otp:attempts:%s", in.Email)
	attempts, err := u.cache.Incr(ctx, attemptsKey, dotp.AttemptsWindow)
	if err != nil {
		return err
	}
	if attempts > dotp.MaxAttempts {
		return ucAuth.ErrTooManyAttempts
	}

	// validate otp
	otpData, err := u.otpRepo.Find(ctx, in.Email, in.Code)
	if err != nil {
		return err
	}
	if otpData == nil || !otpData.IsValid() {
		return ucAuth.ErrInvalidOTP
	}

	// validate and get corporate data
	corporateData, err := u.corporateRepo.FindCorporate(ctx, uuid.Nil, in.CorporateAccountNumber)
	if err != nil {
		return err
	}
	if corporateData == nil {
		return ucAuth.ErrCorporateNotFound
	}

	// validate is user exist
	users, err := u.userRepo.FindUsers(ctx, &in.Username, &in.Email, nil)
	if err != nil {
		return err
	}
	if len(users) > 0 {
		return ucAuth.ErrUserExists
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

	// the code is single-use
	_ = u.cache.Delete(ctx, attemptsKey)
	return u.otpRepo.Deactivate(ctx, otpData.ID)
}
