package dto

import "github.com/google/uuid"

type GetTransactionsResponse struct {
	Data    []DataTransactionsResponse `json:"data"`
	Total   int32                      `json:"total"`
	Page    int                        `json:"page"`
	PerPage int                        `json:"per_page"`
}

type DataTransactionsResponse struct {
	TransactionID       uuid.UUID `json:"id"`
	RefNum              string    `json:"ref_num"`
	TotalTransferAmount float64   `json:"total_transfer_amount"`
	TotalTransferRecord int32     `json:"total_transfer_record"`
	FromAccountNo       string    `json:"from_account_no"`
	Maker               string    `json:"maker"`
	TransferDate        string    `json:"transfer_date"`
	Status              string    `json:"status"`
}
