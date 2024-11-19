package module

import (
	"context"
	"fmt"
	"time"
	dotp "tx-bank/internal/domain/otp"
)

func (u *usecase) SendOTP(ctx context.Context, email string) error {
	var (
		err error
	)

	otp := dotp.NewOTP(email)

	if err = u.otpRepo.Store(ctx, otp); err != nil {
		return err
	}

	// TODO: add retry mechanism when send fail
	if err = u.notificationSvc.SendMail(email,
		"Subject: Your OTP Code\n",
		fmt.Sprintf("Your OTP code is: %s.\nPlease use it before %s.", otp, otp.Expiry().Format(time.RFC1123)),
	); err != nil {
		return err
	}
	return nil
}
