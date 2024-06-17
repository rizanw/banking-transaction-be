package transaction

import "time"

type AuditLogDB struct {
	ID            int64
	TransactionID int64
	UserID        int64
	Action        string
	Timestamp     time.Time
}
