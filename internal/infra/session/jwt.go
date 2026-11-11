package session

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
)

const (
	formatKeyBlacklistedToken = "token:blacklisted:%s"
)

type jwtManager struct {
	secretKey string
	redis     *redis.Client
}

func NewJWTManager(secretKey string, redis *redis.Client) Manager {
	return &jwtManager{
		secretKey: secretKey,
		redis:     redis,
	}
}

func (m *jwtManager) GenerateToken(userID string, role int32, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(m.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (m *jwtManager) ValidateToken(token string) (string, int32, error) {
	// fail closed: if the blacklist can't be checked, reject the token
	blacklisted, err := m.isTokenBlacklisted(token)
	if err != nil {
		return "", 0, fmt.Errorf("check token blacklist: %w", err)
	}
	if blacklisted {
		return "", 0, errors.New("expired token")
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(m.secretKey), nil
	})
	if err != nil || !parsedToken.Valid {
		return "", 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", 0, errors.New("invalid token claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", 0, errors.New("user ID not found in token")
	}
	role, ok := claims["role"].(float64)
	if !ok {
		return "", 0, errors.New("role not found in token")
	}

	return userID, int32(role), nil
}

// BlacklistToken revokes a token until it would have expired anyway.
func (m *jwtManager) BlacklistToken(token string) error {
	return m.redis.Set(context.TODO(), fmt.Sprintf(formatKeyBlacklistedToken, token), true, 25*time.Hour).Err()
}

func (m *jwtManager) isTokenBlacklisted(token string) (bool, error) {
	result, err := m.redis.Exists(context.TODO(), fmt.Sprintf(formatKeyBlacklistedToken, token)).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}
