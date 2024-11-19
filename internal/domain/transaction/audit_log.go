package transaction

import (
	"time"

	"github.com/google/uuid"
)

const (
	AuditActionApprove = "approve"
	AuditActionReject  = "reject"
)

type AuditLog struct {
	ID            uuid.UUID
	TransactionID uuid.UUID
	UserID        uuid.UUID
	Action        string
	Timestamp     time.Time
}

func NewAuditLog(trxID, userID uuid.UUID, action string) *AuditLog {
	return &AuditLog{
		ID:            uuid.New(),
		TransactionID: trxID,
		UserID:        userID,
		Action:        action,
		Timestamp:     time.Now(),
	}
}
