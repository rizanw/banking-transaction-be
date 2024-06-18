package module

import (
	"context"
	"tx-bank/internal/model/transaction"
)

func (u *usecase) GetTransactions(ctx context.Context, in transaction.TransactionRequest) (transaction.TransactionResponse, error) {
	var (
		txData []transaction.Transaction
		filter transaction.TransactionFilter = in.Filter
	)

	if filter.CorporateID > 0 {
		users, err := u.db.FindUsers("", "", 0, filter.CorporateID)
		if err != nil {
			return transaction.TransactionResponse{}, err
		}
		usrs := make([]int64, 0)
		for _, usr := range users {
			usrs = append(usrs, usr.ID)
		}
		filter.Makers = usrs
	}

	txs, total, err := u.db.GetTransactions(filter, (in.Page-1)*in.PerPage, in.PerPage)
	if err != nil {
		return transaction.TransactionResponse{}, err
	}
	if len(txs) == 0 {
		return transaction.TransactionResponse{
			Data:  nil,
			Total: 0,
			Page:  1,
		}, nil
	}

	makers, err := u.db.FindUsers("", "", txs[0].Maker, 0)
	if err != nil || len(makers) == 0 {
		return transaction.TransactionResponse{}, err
	}

	corp, err := u.db.FindCorporate(makers[0].CorporateID, "")
	if err != nil {
		return transaction.TransactionResponse{}, err
	}

	for _, tx := range txs {
		txData = append(txData, transaction.Transaction{
			TransactionID:       tx.ID,
			RefNum:              tx.RefNum,
			TotalTransferAmount: tx.AmountTotal,
			TotalTransferRecord: tx.RecordTotal,
			FromAccountNo:       corp.AccountNum,
			Maker:               makers[0].Username,
			TransferDate:        tx.TxDate.Format(layoutDateTime),
			Status:              tx.GetStatusName(),
		})
	}

	return transaction.TransactionResponse{
		Data:    txData,
		Total:   total,
		Page:    in.Page,
		PerPage: in.PerPage,
	}, nil
}
