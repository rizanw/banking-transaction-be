package transaction

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"
	dtransaction "tx-bank/internal/domain/transaction"
	"tx-bank/internal/dto"

	"github.com/google/uuid"
)

func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		ctx  = r.Context()
		req  dto.UploadTransactionRequest
		resp struct {
			Message string   `json:"message"`
			Errors  []string `json:"errors"`
		}
		csvContent []dto.UploadTransactionCSV
	)

	w.Header().Set("Content-Type", "application/json")

	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		http.Error(w, "user not login", http.StatusUnauthorized)
	}

	req, csvContent, err = parseUploadRequest(r)
	if err != nil {
		resp.Message = err.Error()
		resp.Errors = generateUploadErrorData(csvContent)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	req.MakerID = uuid.MustParse(userID)
	err = h.ucTransaction.UploadTransaction(ctx, req, csvContent)
	if err != nil {
		resp.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp.Message = "success"
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func generateUploadErrorData(txData []dto.UploadTransactionCSV) []string {
	errorData := make([]string, len(txData))
	for _, txDatum := range txData {
		errorData = append(errorData, txDatum.Message)
	}
	return errorData
}

func parseUploadRequest(r *http.Request) (dto.UploadTransactionRequest, []dto.UploadTransactionCSV, error) {
	var (
		err                   error
		csvContent            [][]string
		transactionData       []dto.UploadTransactionCSV
		isInstructionStanding bool
		instructionDate       time.Time
	)

	instructionType := r.FormValue("instruction_type")
	switch instructionType {
	case dtransaction.InstructionTypeImmediate:
		isInstructionStanding = false
	case dtransaction.InstructionTypeStanding:
		isInstructionStanding = true
	default:
		return dto.UploadTransactionRequest{}, nil, errors.New("invalid instruction type")
	}

	if isInstructionStanding {
		instructionDateStr := r.FormValue("transaction_date")
		instructionDateInt, err := strconv.ParseInt(instructionDateStr, 10, 64)
		if err != nil {
			return dto.UploadTransactionRequest{}, nil, err
		}
		instructionDate = time.Unix(instructionDateInt, 0)
	}

	totalRecordStr := r.FormValue("total_record")
	totalRecord, err := strconv.Atoi(totalRecordStr)
	if err != nil {
		return dto.UploadTransactionRequest{}, nil, err
	}

	totalAmountStr := r.FormValue("total_amount")
	totalAmount, err := strconv.ParseFloat(totalAmountStr, 64)
	if err != nil {
		return dto.UploadTransactionRequest{}, nil, err
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		return dto.UploadTransactionRequest{}, nil, fmt.Errorf("fail to read form named file")
	}
	err = checkFileFormat(header, "csv")
	if err != nil {
		return dto.UploadTransactionRequest{}, nil, err
	}
	csvContent, err = parseFileCSV(file, header)
	if err != nil {
		return dto.UploadTransactionRequest{}, nil, err
	}

	isErrorValidate := false
	for idx, row := range csvContent[1:] {
		txRow := dto.UploadTransactionCSV{}

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
		return dto.UploadTransactionRequest{}, transactionData, errors.New("invalid csv content value")
	}

	return dto.UploadTransactionRequest{
		InstructionType: instructionType,
		TransactionDate: instructionDate,
		TotalAmount:     totalAmount,
		TotalRecord:     int32(totalRecord),
	}, transactionData, nil
}

func checkFileFormat(header *multipart.FileHeader, expectedFormat string) error {
	fileDetails := strings.Split(header.Filename, ".")
	if len(fileDetails) < 2 {
		return errors.New("file doesn't have extension")
	}

	if fileDetails[1] != expectedFormat {
		return errors.New("file extension doesn't match with requirement")
	}

	return nil
}

func parseFileCSV(file multipart.File, header *multipart.FileHeader) ([][]string, error) {
	var (
		err        error
		csvContent [][]string
	)

	if err = checkFileFormat(header, "csv"); err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(file)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return csvContent, err
		}

		csvContent = append(csvContent, record)
	}

	return csvContent, nil
}
