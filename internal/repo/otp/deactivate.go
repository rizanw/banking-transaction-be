package otp

import (
	"context"

	"github.com/google/uuid"
)

const qDeactivateOTP = `
	UPDATE "otps" SET "is_active" = false WHERE "id" = $1;
`

// Deactivate marks an OTP as used so it can't be used again.
func (r *repo) Deactivate(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, qDeactivateOTP, id)
	return err
}
