package module

import (
	dcorporate "tx-bank/internal/domain/corporate"
	dotp "tx-bank/internal/domain/otp"
	duser "tx-bank/internal/domain/user"
	"tx-bank/internal/infra/notification"
	"tx-bank/internal/infra/session"
	ucAuth "tx-bank/internal/usecase/auth"
)

type usecase struct {
	otpRepo         dotp.Repository
	userRepo        duser.Repository
	corporateRepo   dcorporate.Repository
	notificationSvc notification.Service
	sessionMgr      session.Manager
}

func New(sessionMgr session.Manager, otpRepo dotp.Repository, userRepo duser.Repository, corporateRepo dcorporate.Repository, notificationSvc notification.Service) ucAuth.UseCase {
	return &usecase{
		sessionMgr:      sessionMgr,
		otpRepo:         otpRepo,
		userRepo:        userRepo,
		corporateRepo:   corporateRepo,
		notificationSvc: notificationSvc,
	}
}
