package auth

import "time"

type OTP struct {
	ID     int64     `db:"id"`
	UserID int64     `db:"user_id"`
	Code   string    `db:"code"`
	Expire time.Time `db:"expires_at"`
}
