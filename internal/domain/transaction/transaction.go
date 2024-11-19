package transaction

import (
	"strings"
	"time"
	"tx-bank/internal/domain/user"

	"github.com/google/uuid"
)

type Transaction struct {
	ID              uuid.UUID
	RefNum          string
	AmountTotal     float64
	RecordTotal     int32
	Maker           *user.User
	TxDate          time.Time
	Status          Status
	InstructionType string
	CreatedAt       time.Time
	Details         []Detail
}

func NewTransaction(maker *user.User, txDate time.Time, status Status, instructionType string, amountTotal float64, recordTotal int32, txDetails ...Detail) *Transaction {
	return &Transaction{
		ID:              uuid.New(),
		RefNum:          strings.ReplaceAll(uuid.NewString(), "-", ""),
		AmountTotal:     amountTotal,
		RecordTotal:     recordTotal,
		Maker:           maker,
		TxDate:          txDate,
		Status:          status,
		InstructionType: instructionType,
		CreatedAt:       time.Now(),
		Details:         txDetails,
	}
}

type TransactionFilter struct {
	Status             Status
	Makers             []uuid.UUID
	StartDate, EndDate time.Time
	Offset, Limit      int
}
