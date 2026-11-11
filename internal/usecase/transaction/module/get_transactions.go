package module

import (
	"context"
	domain "tx-bank/internal/domain/transaction"
	"tx-bank/internal/dto"
)

func (u *usecase) GetTransactions(ctx context.Context, in dto.GetTransactionsRequest) (dto.GetTransactionsResponse, error) {
	_, makers, err := u.corporateMakers(ctx, in.RequesterID)
	if err != nil {
		return dto.GetTransactionsResponse{}, err
	}

	filter := domain.TransactionFilter{
		Status:    in.Filter.Status,
		Makers:    makers,
		StartDate: in.Filter.StartDate,
		EndDate:   in.Filter.EndDate,
		Offset:    (in.Page - 1) * in.PerPage,
		Limit:     in.PerPage,
	}

	total, transactions, err := u.transactionRepo.GetTransactions(ctx, filter)
	if err != nil {
		return dto.GetTransactionsResponse{}, err
	}

	transactionsData := make([]dto.DataTransactionsResponse, 0, len(transactions))
	for _, tx := range transactions {
		maker, err := u.userRepo.GetUserByID(ctx, tx.Maker.ID)
		if err != nil {
			return dto.GetTransactionsResponse{}, err
		}
		if maker == nil {
			continue
		}
		corporate, err := u.corporateRepo.GetCorporateByID(ctx, maker.Corporate.ID)
		if err != nil {
			return dto.GetTransactionsResponse{}, err
		}

		transactionsData = append(transactionsData, dto.DataTransactionsResponse{
			TransactionID:       tx.ID,
			RefNum:              tx.RefNum,
			TotalTransferAmount: tx.AmountTotal,
			TotalTransferRecord: tx.RecordTotal,
			FromAccountNo:       corporate.AccountNum,
			Maker:               maker.Username,
			TransferDate:        tx.TxDate.Format(layoutDateTime),
			Status:              tx.Status.GetName(),
		})
	}

	return dto.GetTransactionsResponse{
		Data:    transactionsData,
		Total:   total,
		Page:    in.Page,
		PerPage: in.PerPage,
	}, nil
}
