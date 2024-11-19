package otp

import (
	"context"
	domain "tx-bank/internal/domain/otp"
)

const qInsertOTP = `
	INSERT INTO "otps"
		("recipient", "code")
	VALUES 
		($1,$2);
`

func (r *repo) Store(ctx context.Context, otp *domain.OTP) error {
	_, err := r.db.ExecContext(ctx, qInsertOTP, otp.Recipient, otp.Code)
	if err != nil {
		return err
	}

	return nil
}
