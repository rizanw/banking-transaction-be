package module

import (
	"context"
	domain "tx-bank/internal/domain/transaction"
	"tx-bank/internal/dto"

	"github.com/google/uuid"
)

func (u *usecase) GetTransactions(ctx context.Context, in dto.GetTransactionsRequest) (dto.GetTransactionsResponse, error) {
	filter := domain.TransactionFilter{
		Status:    in.Filter.Status,
		StartDate: in.Filter.StartDate,
		EndDate:   in.Filter.EndDate,
		Offset:    (in.Page - 1) * in.PerPage,
		Limit:     in.PerPage,
	}
	if len(in.Filter.CorporateID.String()) > 0 {
		users, err := u.userRepo.FindUsers(ctx, nil, nil, &in.Filter.CorporateID)
		if err != nil {
			return dto.GetTransactionsResponse{}, err
		}
		makers := make([]uuid.UUID, 0)
		for _, user := range users {
			makers = append(makers, user.ID)
		}
		filter.Makers = makers
	}

	total, transactions, err := u.transactionRepo.GetTransactions(ctx, filter)
	if err != nil {
		return dto.GetTransactionsResponse{}, err
	}
	if len(transactions) == 0 {
		return dto.GetTransactionsResponse{
			Data:    []dto.DataTransactionsResponse{},
			Total:   total,
			Page:    in.Page,
			PerPage: in.PerPage,
		}, nil
	}

	transactionsData := make([]dto.DataTransactionsResponse, 0)
	for _, tx := range transactions {
		maker, err := u.userRepo.GetUserByID(ctx, tx.Maker.ID)
		if err != nil {
			return dto.GetTransactionsResponse{}, err
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
