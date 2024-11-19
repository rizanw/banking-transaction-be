package notification

type Service interface {
	SendMail(receiver, subject, content string) error
}
