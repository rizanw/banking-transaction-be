package dto

import (
	"time"
	"tx-bank/internal/domain/transaction"

	"github.com/google/uuid"
)

type GetTransactionsRequest struct {
	RequesterID uuid.UUID // logged-in user; results are limited to their corporate
	Page        int
	PerPage     int
	Filter      FilterTransactionsRequest
}

type FilterTransactionsRequest struct {
	Status    transaction.Status
	StartDate time.Time
	EndDate   time.Time
}
