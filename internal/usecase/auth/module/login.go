package module

import (
	"context"
	"errors"
	"time"
	"tx-bank/internal/dto"

	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) Login(ctx context.Context, in dto.LoginRequest) (dto.LoginResponse, error) {

	// find existing user
	userData, err := u.userRepo.GetUserByUsername(ctx, in.Username)
	if err != nil {
		return dto.LoginResponse{}, err
	}
	if userData == nil {
		return dto.LoginResponse{}, errors.New("user not found")
	}

	// compare hashed password
	if err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(in.Password)); err != nil {
		return dto.LoginResponse{}, errors.New("invalid password")
	}

	// find corp data
	corpData, err := u.corporateRepo.FindCorporate(ctx, userData.Corporate.ID, "")
	if err != nil {
		return dto.LoginResponse{}, err
	}

	// generate token
	token, err := u.sessionMgr.GenerateToken(userData.ID.String(), int32(userData.Role), 24*time.Hour)
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
