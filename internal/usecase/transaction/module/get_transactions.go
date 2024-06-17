package module

import (
	"time"
	"tx-bank/internal/model/transaction"
)

func (u *usecase) GetTransactions(in transaction.TransactionRequest) (transaction.TransactionResponse, error) {
	var (
		txData []transaction.Transaction
	)

	txs, total, err := u.db.GetTransactions(in.Page, in.PerPage)
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

	maker, err := u.db.FindUser("", "", txs[0].Maker)
	if err != nil {
		return transaction.TransactionResponse{}, err
	}

	corp, err := u.db.FindCorporate(maker.CorporateID, "")
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
			Maker:               maker.Username,
			TransferDate:        tx.TxDate.Format(time.RFC850),
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
