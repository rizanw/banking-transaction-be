package module

import (
	"context"
	"fmt"
	"log"
	"time"
	dotp "tx-bank/internal/domain/otp"
)

func (u *usecase) SendOTP(ctx context.Context, email string) error {
	otp := dotp.NewOTP(email)
	if err := u.otpRepo.Store(ctx, otp); err != nil {
		return err
	}

	// local development without SMTP credentials: print the code instead of emailing it
	if !u.notificationSvc.Enabled() {
		log.Printf("[otp] SMTP is not configured; OTP for %s is %s", email, otp.Code)
		return nil
	}

	// TODO: add retry mechanism when send fail
	return u.notificationSvc.SendMail(email,
		"Subject: Your OTP Code\n",
		fmt.Sprintf("Your OTP code is: %s.\nPlease use it before %s.", otp.Code, otp.Expiry().Format(time.RFC1123)),
	)
}
