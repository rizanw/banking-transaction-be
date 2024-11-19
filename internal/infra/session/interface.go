package session

import "time"

type Manager interface {
	GenerateToken(userID string, role int32, duration time.Duration) (string, error)
	ValidateToken(token string) (string, int32, error)
	BlacklistToken(token string) error
}
