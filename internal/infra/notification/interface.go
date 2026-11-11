package notification

type Service interface {
	// Enabled reports whether the service is configured to send emails.
	Enabled() bool
	SendMail(receiver, subject, content string) error
}
