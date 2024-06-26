package user

type UserDB struct {
	ID          int64  `db:"id"`
	Username    string `db:"username"`
	Password    string `db:"password"`
	Email       string `db:"email"`
	Phone       string `db:"phone"`
	CorporateID int64  `db:"corporate_id"`
	Role        int32  `db:"role"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     int32  `json:"role"`
}
