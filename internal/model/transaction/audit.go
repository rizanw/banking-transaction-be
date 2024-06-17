package transaction

import "time"

const (
	AuditActionApprove = "approve"
	AuditActionReject  = "reject"
)

type AuditLogDB struct {
	ID            int64
	TransactionID int64
	UserID        int64
	Action        string
	Timestamp     time.Time
}
