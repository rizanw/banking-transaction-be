package transaction

import "time"

const (
	StatusNone = iota
	StatusAwaitingApproval
	StatusApproved
	StatusRejected
)

type TransactionDB struct {
	ID              int64     `db:"id"`
	RefNum          string    `db:"ref_num"`
	AmountTotal     float64   `db:"amount_total"`
	RecordTotal     int32     `db:"record_total"`
	Maker           int64     `db:"maker"`
	TxDate          time.Time `db:"date"`
	Status          int32     `db:"status"`
	InstructionType string    `db:"instruction_type"`
	CreatedAt       time.Time `db:"created_at"`
}

func (t TransactionDB) GetStatusName() string {
	switch t.Status {
	case StatusAwaitingApproval:
		return "Awaiting Approval"
	case StatusApproved:
		return "Approved"
	case StatusRejected:
		return "Rejected"
	}
	return ""
}

type TransactionFilter struct {
	Status      int
	Makers      []int64
	CorporateID int64
	StartDate   time.Time
	EndDate     time.Time
}

type TransactionRequest struct {
	TransactionID int64
	Page          int
	PerPage       int
	Filter        TransactionFilter
}

type TransactionResponse struct {
	Data    []Transaction `json:"data"`
	Total   int32         `json:"total"`
	Page    int           `json:"page"`
	PerPage int           `json:"per_page"`
}

type Transaction struct {
	TransactionID       int64   `json:"id"`
	RefNum              string  `json:"ref_num"`
	TotalTransferAmount float64 `json:"total_transfer_amount"`
	TotalTransferRecord int32   `json:"total_transfer_record"`
	FromAccountNo       string  `json:"from_account_no"`
	Maker               string  `json:"maker"`
	TransferDate        string  `json:"transfer_date"`
	Status              string  `json:"status"`
}
