package module

import (
	"context"
	dcorporate "tx-bank/internal/domain/corporate"
	dtransaction "tx-bank/internal/domain/transaction"
	duser "tx-bank/internal/domain/user"
	"tx-bank/internal/dto"

	"golang.org/x/sync/errgroup"
)

func (u *usecase) GetTransaction(ctx context.Context, in dto.GetTransactionRequest) (dto.GetTransactionResponse, error) {
	transaction, err := u.transactionRepo.GetTransactionByID(ctx, in.TransactionID)
	if err != nil {
		return dto.GetTransactionResponse{}, err
	}
	if transaction == nil {
		return dto.GetTransactionResponse{}, dtransaction.ErrNotFound
	}

	// hide transactions from other corporates as if they didn't exist
	allowed, err := u.sameCorporate(ctx, in.RequesterID, transaction.Maker.ID)
	if err != nil {
		return dto.GetTransactionResponse{}, err
	}
	if !allowed {
		return dto.GetTransactionResponse{}, dtransaction.ErrNotFound
	}

	var (
		eg                 errgroup.Group
		total              int32
		maker              *duser.User
		corporate          *dcorporate.Corporate
		transactionDetails = make([]dto.DetailTransactionResponse, 0)
	)

	eg.Go(func() error {
		var err error
		maker, err = u.userRepo.GetUserByID(ctx, transaction.Maker.ID)
		if err != nil {
			return err
		}
		if maker == nil {
			return dtransaction.ErrNotFound
		}
		corporate, err = u.corporateRepo.GetCorporateByID(ctx, maker.Corporate.ID)
		return err
	})

	eg.Go(func() error {
		var (
			details []dtransaction.Detail
			err     error
		)
		total, details, err = u.transactionRepo.GetTransactionDetails(ctx, transaction.ID)
		if err != nil {
			return err
		}
		for _, d := range details {
			transactionDetails = append(transactionDetails, dto.DetailTransactionResponse{
				ID:              d.ID,
				ToAccountNumber: d.ToAccountNumber,
				ToAccountName:   d.ToAccountName,
				ToAccountBank:   d.ToAccountBank,
				Amount:          d.Amount,
				Description:     d.Description,
				Status:          d.Status.GetName(),
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
