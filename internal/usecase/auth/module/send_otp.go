package module

import (
	"errors"
	"fmt"
	"time"
	"tx-bank/internal/model/auth"
)

func (u *usecase) SendOTP(email string) error {
	var (
		err    error
		expiry = time.Now().Add(time.Hour * 24)
	)

	if users, err := u.db.FindUsers("", email, 0, 0); err != nil {
		return err
	} else if len(users) > 0 {
		return errors.New("email already in use")
	}

	otp := generateOTP()
	if err = u.sendMail(email,
		"Subject: Your OTP Code\n",
		fmt.Sprintf("Your OTP code is: %s.\nPlease use it before %s.", otp, expiry.Format(time.RFC1123)),
	); err != nil {
		return err
	}

	if err = u.db.InsertOTP(auth.OTP{
		Email:  email,
		Code:   otp,
		Expire: expiry,
	}); err != nil {
		return err
	}

	return nil
}
