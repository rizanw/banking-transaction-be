package session

import "time"

var Sessions = map[string]Session{}

type Session struct {
	UserID int64
	Email  string
	Expiry int64
}

func (s Session) IsExpired() bool {
	expiry := time.Unix(s.Expiry, 0)
	return expiry.Before(time.Now())
}
