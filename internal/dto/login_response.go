package dto

import "github.com/google/uuid"

type LoginResponse struct {
	AccessToken   string                 `json:"access_token"`
	CorporateInfo CorporateLoginResponse `json:"corporate_info"`
	UserInfo      UserLoginResponse      `json:"user_info"`
}

type UserLoginResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Role     string    `json:"role"`
}

type CorporateLoginResponse struct {
	Name       string `json:"name"`
	AccountNum string `json:"account_num"`
}
