package dto

import (
	"time"
	"tx-bank/internal/domain/transaction"

	"github.com/google/uuid"
)

type GetTransactionsRequest struct {
	Page    int
	PerPage int
	Filter  FilterTransactionsRequest
}

type FilterTransactionsRequest struct {
	Status      transaction.Status
	CorporateID uuid.UUID
	StartDate   time.Time
	EndDate     time.Time
}
