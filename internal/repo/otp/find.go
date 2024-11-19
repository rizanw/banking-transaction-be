package otp

import (
	"context"
	"database/sql"
	"errors"
	"time"
	domain "tx-bank/internal/domain/otp"

	"github.com/google/uuid"
)

const qFindOTP = `
	SELECT 
	    "id", "recipient", "code", "is_active", "created_at"
	FROM "otps"
	WHERE
	    recipient = $1 AND code = $2;
`

type otpSchema struct {
	ID        uuid.UUID `db:"id"`
	Recipient string    `db:"recipient"`
	Code      string    `db:"code"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
}

func (r *repo) Find(ctx context.Context, recipient, code string) (*domain.OTP, error) {
	var otp otpSchema

	err := r.db.QueryRowContext(ctx, qFindOTP, recipient, code).Scan(&otp.ID, &otp.Recipient, &otp.Code, &otp.IsActive, &otp.CreatedAt)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return &domain.OTP{
		ID:        otp.ID,
		Recipient: otp.Recipient,
		Code:      otp.Code,
		IsActive:  otp.IsActive,
		CreatedAt: otp.CreatedAt,
	}, nil
}
