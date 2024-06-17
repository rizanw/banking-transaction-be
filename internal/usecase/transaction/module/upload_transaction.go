package module

import (
	"errors"
	"strings"
	"time"
	"tx-bank/internal/model/transaction"

	"github.com/google/uuid"
)

func (u *usecase) UploadTransaction(in transaction.UploadTransactionRequest, csv []transaction.UploadTransactionCSV) error {
	var (
		csvTotalRecords int32
		csvTotalAmount  float64
		transactionDate time.Time
		status          int32 = transaction.StatusAwaitingApproval
		transacDetails        = make([]transaction.TransactionDetailDB, 0)
	)

	// recalculate & validate request
	for _, transac := range csv {
		csvTotalAmount += transac.TransferAmount
		csvTotalRecords++
	}
	if csvTotalRecords != in.TotalRecord {
		return errors.New("total records don't match")
	}
	if csvTotalAmount != in.TotalAmount {
		return errors.New("total amount don't match")
	}

	switch in.InstructionType {
	case transaction.InstructionTypeImmediate:
		transactionDate = time.Now()
	case transaction.InstructionTypeStanding:
		transactionDate = in.TransactionDate

	}

	// insert transaction
	txID, err := u.db.InsertTransaction(transaction.TransactionDB{
		RefNum:          strings.ReplaceAll(uuid.New().String(), "-", ""),
		AmountTotal:     csvTotalAmount,
		RecordTotal:     csvTotalRecords,
		Maker:           in.MakerID,
		TxDate:          transactionDate,
		Status:          status,
		InstructionType: in.InstructionType,
	})
	if err != nil {
		return err
	}

	// bulk insert transaction details
	for _, transac := range csv {
		transacDetails = append(transacDetails, transaction.TransactionDetailDB{
			TransactionID:   txID,
			ToAccountNumber: transac.ToAccountNum,
			ToAccountName:   transac.ToAccountName,
			ToAccountBank:   transac.ToBankName,
			Amount:          transac.TransferAmount,
			Status:          status,
		})
	}
	err = u.db.InsertTransactionDetails(transacDetails)
	if err != nil {
		return err
	}

	return nil
}
