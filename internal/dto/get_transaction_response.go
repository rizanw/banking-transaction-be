package dto

import "github.com/google/uuid"

type GetTransactionResponse struct {
	ID              uuid.UUID                   `json:"id"`
	RefNum          string                      `json:"ref_num"`
	FromAccountNum  string                      `json:"from_account_num"`
	SubmitDateTime  string                      `json:"submit_datetime"`
	TransferDate    string                      `json:"transfer_date"`
	InstructionType string                      `json:"instruction_type"`
	Maker           string                      `json:"maker"`
	TotalAmount     float64                     `json:"total_amount"`
	TotalRecord     int32                       `json:"total_record"`
	Data            []DetailTransactionResponse `json:"data"`
	Total           int32                       `json:"total"`
	Page            int                         `json:"page"`
	PerPage         int                         `json:"per_page"`
}

type DetailTransactionResponse struct {
	ID              uuid.UUID `json:"id"`
	ToAccountNumber string    `json:"to_account_num"`
	ToAccountName   string    `json:"to_account_name"`
	ToAccountBank   string    `json:"to_account_bank"`
	Amount          float64   `json:"amount"`
	Description     string    `json:"description"`
	Status          string    `json:"status"`
}
