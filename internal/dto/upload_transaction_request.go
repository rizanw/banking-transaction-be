package dto

import (
	"time"

	"github.com/google/uuid"
)

type UploadTransactionRequest struct {
	MakerID         uuid.UUID
	InstructionType string    `json:"instruction_type"`
	TransactionDate time.Time `json:"transaction_date"`
	TotalAmount     float64   `json:"total_amount"`
	TotalRecord     int32     `json:"total_record"`
}

type UploadTransactionCSV struct {
	Message        string
	ToBankName     string
	ToAccountNum   string
	ToAccountName  string
	TransferAmount float64
}
