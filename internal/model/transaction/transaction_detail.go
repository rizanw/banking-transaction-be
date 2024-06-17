package transaction

type TransactionDetailDB struct {
	ID              int64   `db:"id"`
	TransactionID   int64   `db:"transaction_id"`
	ToAccountNumber string  `db:"to_account_num"`
	ToAccountName   string  `db:"to_account_name"`
	ToAccountBank   string  `db:"to_account_bank"`
	Amount          float64 `db:"amount"`
	Description     string  `db:"description"`
	Status          int32   `db:"status"`
}

func (t TransactionDetailDB) GetStatusName() string {
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

type TransactionDetailResponse struct {
	ID              int64               `json:"id"`
	RefNum          string              `json:"ref_num"`
	FromAccountNum  string              `json:"from_account_num"`
	SubmitDateTime  string              `json:"submit_datetime"`
	TransferDate    string              `json:"transfer_date"`
	InstructionType string              `json:"instruction_type"`
	Maker           string              `json:"maker"`
	TotalAmount     float64             `json:"total_amount"`
	TotalRecord     int32               `json:"total_record"`
	Data            []TransactionDetail `json:"data"`
	Total           int32               `json:"total"`
	Page            int                 `json:"page"`
	PerPage         int                 `json:"per_page"`
}

type TransactionDetail struct {
	ID              int64   `json:"id"`
	ToAccountNumber string  `json:"to_account_num"`
	ToAccountName   string  `json:"to_account_name"`
	ToAccountBank   string  `json:"to_account_bank"`
	Amount          float64 `json:"amount"`
	Description     string  `json:"description"`
	Status          string  `json:"status"`
}
