package dto

import "github.com/google/uuid"

type GetTransactionRequest struct {
	TransactionID uuid.UUID
	Page          int
	PerPage       int
}
