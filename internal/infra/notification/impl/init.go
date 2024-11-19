package impl

import (
	"tx-bank/internal/config"
	"tx-bank/internal/infra/notification"
)

type NotificationService struct {
	SMTPConf config.SMTPConfig
}

func New(conf config.SMTPConfig) notification.Service {
	return &NotificationService{
		SMTPConf: conf,
	}
}
