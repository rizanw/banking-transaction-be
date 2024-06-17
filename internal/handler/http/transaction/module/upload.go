package module

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"tx-bank/internal/common/session"
	"tx-bank/internal/model/transaction"
)

func (h *handler) Upload(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		ctx        = r.Context()
		request    transaction.UploadTransactionRequest
		csvContent []transaction.UploadTransactionCSV
	)

	w.Header().Set("Content-Type", "application/json")

	ses, ok := ctx.Value("session").(session.Session)
	if ses.UserID == 0 || !ok {
		http.Error(w, "user not login", http.StatusUnauthorized)
	}

	request, csvContent, err = parseUploadRequest(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(transaction.UploadTransactionResponse{
			Message: err.Error(),
			Errors:  generateUploadErrorData(csvContent),
		})
		return
	}

	request.MakerID = ses.UserID
	request.MakerRole = ses.Role
	err = h.ucTransaction.UploadTransaction(request, csvContent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(transaction.UploadTransactionResponse{
			Message: err.Error(),
			Errors:  nil,
		})
		return
	}

	err = json.NewEncoder(w).Encode(transaction.UploadTransactionResponse{
		Message: "success",
		Errors:  nil,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func generateUploadErrorData(txData []transaction.UploadTransactionCSV) []string {
	errorData := make([]string, len(txData))
	for _, txDatum := range txData {
		errorData = append(errorData, txDatum.Message)
	}
	return errorData
}

func parseUploadRequest(r *http.Request) (transaction.UploadTransactionRequest, []transaction.UploadTransactionCSV, error) {
	var (
		err                   error
		csvContent            [][]string
		transactionData       []transaction.UploadTransactionCSV
		isInstructionStanding bool
		instructionDate       time.Time
	)

	instructionType := r.FormValue("instruction_type")
	switch instructionType {
	case transaction.InstructionTypeImmediate:
		isInstructionStanding = false
	case transaction.InstructionTypeStanding:
		isInstructionStanding = true
	default:
		return transaction.UploadTransactionRequest{}, nil, errors.New("invalid instruction type")
	}

	if isInstructionStanding {
		instructionDateStr := r.FormValue("transaction_date")
		instructionDateInt, err := strconv.ParseInt(instructionDateStr, 10, 64)
		if err != nil {
			return transaction.UploadTransactionRequest{}, nil, err
		}
		instructionDate = time.Unix(instructionDateInt, 0)
	}

	totalRecordStr := r.FormValue("total_record")
	totalRecord, err := strconv.Atoi(totalRecordStr)
	if err != nil {
		return transaction.UploadTransactionRequest{}, nil, err
	}

	totalAmountStr := r.FormValue("total_amount")
	totalAmount, err := strconv.ParseFloat(totalAmountStr, 64)
	if err != nil {
		return transaction.UploadTransactionRequest{}, nil, err
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		return transaction.UploadTransactionRequest{}, nil, fmt.Errorf("fail to read form named file")
	}
	err = checkFileFormat(header, "csv")
	if err != nil {
		return transaction.UploadTransactionRequest{}, nil, err
	}
	csvContent, err = parseFileCSV(file, header)
	if err != nil {
		return transaction.UploadTransactionRequest{}, nil, err
	}
	isErrorValidate := false
	for idx, row := range csvContent[1:] {
		txRow := transaction.UploadTransactionCSV{}

		if len(row) != 4 {
			isErrorValidate = true
			txRow.Message = fmt.Sprintf("[row:%d] invalid column length", idx-1)
			continue
		}

		toBankName := row[0]
		if toBankName == "" {
			isErrorValidate = true
			txRow.Message = fmt.Sprintf("[row:%d] invalid to_bank_name format", idx-1)
			continue
		} else {
			txRow.ToBankName = toBankName
		}

		toAccountNum := row[1]
		if toAccountNum == "" {
			isErrorValidate = true
			txRow.Message = fmt.Sprintf("[row:%d] invalid to_account_no format", idx-1)
			continue
		} else {
			txRow.ToAccountNum = toAccountNum
		}

		toAccountName := row[2]
		if toAccountName == "" {
			isErrorValidate = true
			txRow.Message = fmt.Sprintf("[row:%d] invalid to_account_name format", idx-1)
			continue
		} else {
			txRow.ToAccountName = toAccountName
		}

		transferAmount, err := strconv.ParseFloat(row[3], 64)
		if err != nil || transferAmount <= 0 {
			isErrorValidate = true
			txRow.Message = fmt.Sprintf("[row:%d] invalid total_amount format", idx-1)
			continue
		} else {
			txRow.TransferAmount = transferAmount
		}

		transactionData = append(transactionData, txRow)
	}
	if isErrorValidate {
		return transaction.UploadTransactionRequest{}, transactionData, errors.New("invalid csv content value")
	}

	return transaction.UploadTransactionRequest{
		InstructionType: instructionType,
		TransactionDate: instructionDate,
		TotalAmount:     totalAmount,
		TotalRecord:     int32(totalRecord),
	}, transactionData, nil
}
