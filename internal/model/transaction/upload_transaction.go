package transaction

import "time"

const (
	InstructionTypeImmediate = "immediate"
	InstructionTypeStanding  = "standing"
)

type UploadTransactionRequest struct {
	MakerID         int64
	MakerRole       int32
	InstructionType string    `json:"instruction_type"`
	TransactionDate time.Time `json:"transaction_date"`
	TotalAmount     float64   `json:"total_amount"`
	TotalRecord     int32     `json:"total_record"`
}

type UploadTransactionResponse struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

type UploadTransactionCSV struct {
	Message        string
	ToBankName     string
	ToAccountNum   string
	ToAccountName  string
	TransferAmount float64
}
