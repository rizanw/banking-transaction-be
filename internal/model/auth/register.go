package auth

type RegisterRequest struct {
	Username               string `json:"username"`
	Password               string `json:"password"`
	Email                  string `json:"email"`
	Phone                  string `json:"phone"`
	Role                   string `json:"role"`
	CorporateAccountNumber string `json:"corporate_account_number"`
	CorporateName          string `json:"corporate_name"`
}
