package dto

import "github.com/google/uuid"

type GetTransactionRequest struct {
	RequesterID   uuid.UUID // logged-in user; must belong to the maker's corporate
	TransactionID uuid.UUID
	Page          int
	PerPage       int
}
