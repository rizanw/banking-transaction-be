package transaction

import "time"

const (
	StatusNone = iota
	StatusAwaitingApproval
	StatusApproved
	StatusRejected
)

type TransactionDB struct {
	ID          int64     `db:"id"`
	RefNum      string    `db:"ref_num"`
	AmountTotal float64   `db:"amount_total"`
	RecordTotal int32     `db:"record_total"`
	Maker       int64     `db:"maker"`
	TxDate      time.Time `db:"date"`
	Status      int32     `db:"status"`
}
