package module

import (
	"context"
	"errors"
	dcorporate "tx-bank/internal/domain/corporate"
	dtransaction "tx-bank/internal/domain/transaction"
	duser "tx-bank/internal/domain/user"
	"tx-bank/internal/dto"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

func (u *usecase) GetTransaction(ctx context.Context, in dto.GetTransactionRequest) (dto.GetTransactionResponse, error) {
	var (
		eg        errgroup.Group
		err       error
		total     int32
		maker     *duser.User
		corporate *dcorporate.Corporate
	)

	transaction, err := u.transactionRepo.GetTransactionByID(ctx, in.TransactionID)
	if err != nil {
		return dto.GetTransactionResponse{}, err
	}
	if transaction == nil || transaction.ID == uuid.Nil {
		return dto.GetTransactionResponse{}, errors.New("no transaction found")
	}

	eg.Go(func() error {
		maker, err = u.userRepo.GetUserByID(ctx, transaction.Maker.ID)
		if err != nil {
			return err
		}
		corporate, err = u.corporateRepo.GetCorporateByID(ctx, maker.Corporate.ID)
		if err != nil {
			return err
		}
		return nil
	})

	transactionDetails := make([]dto.DetailTransactionResponse, 0)
	eg.Go(func() error {
		trxDetails := make([]dtransaction.Detail, 0)
		total, trxDetails, err = u.transactionRepo.GetTransactionDetails(ctx, transaction.ID)
		if err != nil {
			return err
		}

		for _, trx := range trxDetails {
			transactionDetails = append(transactionDetails, dto.DetailTransactionResponse{
				ID:              trx.ID,
				ToAccountNumber: trx.ToAccountNumber,
				ToAccountName:   trx.ToAccountName,
				ToAccountBank:   trx.ToAccountBank,
				Amount:          trx.Amount,
				Description:     trx.Description,
				Status:          trx.Status.GetName(),
			})
		}
		return nil
	})

	if err = eg.Wait(); err != nil {
		return dto.GetTransactionResponse{}, err
	}

	return dto.GetTransactionResponse{
		ID:              transaction.ID,
		RefNum:          transaction.RefNum,
		FromAccountNum:  corporate.AccountNum,
		SubmitDateTime:  transaction.CreatedAt.Format(layoutDateTime),
		TransferDate:    transaction.TxDate.Format(layoutDateTime),
		InstructionType: transaction.InstructionType,
		Maker:           maker.Username,
		TotalAmount:     transaction.AmountTotal,
		TotalRecord:     transaction.RecordTotal,
		Data:            transactionDetails,
		Total:           total,
		Page:            in.Page,
		PerPage:         in.PerPage,
	}, nil
}
