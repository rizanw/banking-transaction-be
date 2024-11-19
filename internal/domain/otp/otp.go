package otp

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/google/uuid"
)

const (
	TTL = 24 * time.Hour
)

type OTP struct {
	ID        uuid.UUID
	Recipient string
	Code      string
	TTL       time.Duration
	IsActive  bool
	CreatedAt time.Time
}

func NewOTP(recipient string) *OTP {
	return &OTP{
		ID:        uuid.New(),
		Recipient: recipient,
		Code:      generateOTP(),
		IsActive:  true,
		TTL:       TTL,
		CreatedAt: time.Now(),
	}
}

func (o *OTP) IsValid() bool {
	if o.TTL.Milliseconds() == 0 {
		o.TTL = TTL
	}
	return o.IsActive && time.Since(o.CreatedAt) < o.TTL
}

func (o *OTP) Expiry() time.Time {
	return o.CreatedAt.Add(o.TTL)
}

func generateOTP() string {
	// Generate a random 6-digit number securely using crypto/rand
	n, _ := rand.Int(rand.Reader, big.NewInt(1000000)) // Ignore the error
	// Return a zero-padded 6-digit string
	return fmt.Sprintf("%06d", n)
}
