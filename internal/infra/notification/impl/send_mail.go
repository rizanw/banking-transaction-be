package impl

import (
	"fmt"
	"net/smtp"
)

func (s *NotificationService) SendMail(receiver, subject, content string) error {
	from := s.SMTPConf.Username
	to := []string{receiver}
	msg := []byte(subject + "\n" + content)
	smtpServer := fmt.Sprintf("%s:%d", s.SMTPConf.Host, s.SMTPConf.Port)

	err := smtp.SendMail(smtpServer, smtp.PlainAuth(
		"",
		s.SMTPConf.Username,
		s.SMTPConf.Password,
		s.SMTPConf.Host,
	), from, to, msg)
	if err != nil {
		return err
	}

	return nil
}
