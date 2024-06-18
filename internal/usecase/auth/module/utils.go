package module

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"time"
	"tx-bank/internal/common/session"

	"github.com/golang-jwt/jwt"
)

func (u *usecase) generateToken(s session.Session) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": s.UserID,
		"email":   s.Email,
		"role":    s.Role,
		"exp":     s.Expiry,
	})

	signedToken, err := token.SignedString(u.confJWT.Secret)
	if err != nil {
		return "", err
	}

	session.Sessions[signedToken] = s
	return signedToken, nil
}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(1000000)
	return fmt.Sprintf("%06d", randomNumber)
}

func (u *usecase) sendMail(receiver, subject, content string) error {
	auth := smtp.PlainAuth(
		"",
		u.confSMTP.Username,
		u.confSMTP.Password,
		u.confSMTP.Host,
	)

	from := u.confSMTP.Username
	to := []string{receiver}
	msg := []byte(subject + "\n" + content)

	smtpServer := fmt.Sprintf("%s:%d", u.confSMTP.Host, u.confSMTP.Port)
	err := smtp.SendMail(smtpServer, auth, from, to, msg)
	if err != nil {
		return err
	}

	return nil
}
