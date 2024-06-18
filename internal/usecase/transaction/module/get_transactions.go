package module

import (
	"context"
	"tx-bank/internal/common/session"
	"tx-bank/internal/model/transaction"
	"tx-bank/internal/model/user"
)

func (u *usecase) GetTransactions(ctx context.Context, in transaction.TransactionRequest) (transaction.TransactionResponse, error) {
	var (
		ses              = ctx.Value("session").(session.Session)
		filterStatus int = -1
		txData       []transaction.Transaction
	)

	if ses.Role == user.RoleApprover {
		filterStatus = transaction.StatusAwaitingApproval
	}

	txs, total, err := u.db.GetTransactions(filterStatus, (in.Page-1)*in.PerPage, in.PerPage)
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
