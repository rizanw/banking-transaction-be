package auth

import "time"

type OTP struct {
	ID     int64     `db:"id"`
	Email  string    `db:"email"`
	Code   string    `db:"code"`
	Expire time.Time `db:"expires_at"`
}
