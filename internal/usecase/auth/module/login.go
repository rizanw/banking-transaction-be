package module

import (
	"context"
	"time"
	"tx-bank/internal/dto"
	ucAuth "tx-bank/internal/usecase/auth"

	"golang.org/x/crypto/bcrypt"
)

const tokenTTL = 24 * time.Hour

func (u *usecase) Login(ctx context.Context, in dto.LoginRequest) (dto.LoginResponse, error) {
	// find existing user
	userData, err := u.userRepo.GetUserByUsername(ctx, in.Username)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	// same error for unknown user and wrong password, so usernames can't be enumerated
	if userData == nil {
		return dto.LoginResponse{}, ucAuth.ErrInvalidCredentials
	}
	if err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(in.Password)); err != nil {
		return dto.LoginResponse{}, ucAuth.ErrInvalidCredentials
	}

	// find corp data
	corpData, err := u.corporateRepo.FindCorporate(ctx, userData.Corporate.ID, "")
	if err != nil {
		return dto.LoginResponse{}, err
	}
	if corpData == nil {
		return dto.LoginResponse{}, ucAuth.ErrCorporateNotFound
	}

	// generate token
	token, err := u.sessionMgr.GenerateToken(userData.ID.String(), int32(userData.Role), tokenTTL)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		AccessToken: token,
		CorporateInfo: dto.CorporateLoginResponse{
			Name:       corpData.Name,
			AccountNum: corpData.AccountNum,
		},
		UserInfo: dto.UserLoginResponse{
			ID:       userData.ID,
			Username: userData.Username,
			Email:    userData.Email,
			Phone:    userData.Phone,
			Role:     userData.Role.GetName(),
		},
	}, nil
}
