package module

import (
	"time"
	"tx-bank/internal/model/transaction"
)

func (u *usecase) GetTransaction(in transaction.TransactionRequest) (transaction.TransactionDetailResponse, error) {
	var (
		err        error
		trxDetails []transaction.TransactionDetail
	)

	transac, err := u.db.FindTransaction(in.TransactionID)
	if err != nil {
		return transaction.TransactionDetailResponse{}, err
	}
	if transac.ID == 0 {
		return transaction.TransactionDetailResponse{}, nil
	}

	maker, err := u.db.FindUser("", "", transac.Maker)
	if err != nil {
		return transaction.TransactionDetailResponse{}, err
	}

	corp, err := u.db.FindCorporate(maker.CorporateID, "")
	if err != nil {
		return transaction.TransactionDetailResponse{}, err
	}

	transacDetails, total, err := u.db.FindTransactionDetails(transac.ID)
	if err != nil {
		return transaction.TransactionDetailResponse{}, err
	}
	for _, transacDetail := range transacDetails {
		trxDetails = append(trxDetails, transaction.TransactionDetail{
			ID:              transacDetail.ID,
			ToAccountNumber: transacDetail.ToAccountNumber,
			ToAccountName:   transacDetail.ToAccountName,
			ToAccountBank:   transacDetail.ToAccountBank,
			Amount:          transacDetail.Amount,
			Description:     transacDetail.Description,
			Status:          transacDetail.GetStatusName(),
		})
	}

	return transaction.TransactionDetailResponse{
		ID:              in.TransactionID,
		RefNum:          transac.RefNum,
		FromAccountNum:  corp.AccountNum,
		SubmitDateTime:  transac.CreatedAt.Format(time.RFC850),
		TransferDate:    transac.CreatedAt.Format(time.RFC850),
		InstructionType: transac.InstructionType,
		Maker:           maker.Username,
		Data:            trxDetails,
		Total:           total,
		Page:            in.Page,
		PerPage:         in.PerPage,
	}, nil
}
