package module

import (
	"context"
	"errors"
	"time"
	dtransaction "tx-bank/internal/domain/transaction"
	"tx-bank/internal/dto"
)

func (u *usecase) UploadTransaction(ctx context.Context, in dto.UploadTransactionRequest, csv []dto.UploadTransactionCSV) error {
	var (
		csvTotalRecords int32
		csvTotalAmount  float64
		txDate          time.Time
		status          = dtransaction.StatusAwaitingApproval
		txDetails       = make([]dtransaction.Detail, 0)
	)

	// recalculate & validate request
	for _, data := range csv {
		csvTotalAmount += data.TransferAmount
		csvTotalRecords++
		txDetails = append(txDetails, *dtransaction.NewTransactionDetail(
			data.ToAccountNum, data.ToAccountName, data.ToBankName, data.Message, data.TransferAmount, status,
		))
	}
	if csvTotalRecords != in.TotalRecord {
		return errors.New("total records don't match")
	}
	if csvTotalAmount != in.TotalAmount {
		return errors.New("total amount don't match")
	}

	switch in.InstructionType {
	case dtransaction.InstructionTypeImmediate:
		txDate = time.Now()
	case dtransaction.InstructionTypeStanding:
		txDate = in.TransactionDate
	}

	maker, err := u.userRepo.GetUserByID(ctx, in.MakerID)
	if err != nil {
		return err
	}

	// insert transaction & details
	trx := dtransaction.NewTransaction(
		maker, txDate, status, in.InstructionType, csvTotalAmount, csvTotalRecords, txDetails...,
	)
	err = u.transactionRepo.CreateTransaction(ctx, trx)
	if err != nil {
		return err
	}

	return nil
}
