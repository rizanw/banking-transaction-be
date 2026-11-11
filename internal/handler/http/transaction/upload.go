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
	"tx-bank/internal/handler/http/httputil"
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
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	userID, ok := httputil.UserID(r)
	if !ok {
		httputil.Message(w, http.StatusUnauthorized, "user not login")
		return
	}

	req, csvContent, rowErrors, err := parseUploadRequest(r)
	if err != nil {
		resp.Message = err.Error()
		resp.Errors = rowErrors
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	req.MakerID = userID
	err = h.ucTransaction.UploadTransaction(ctx, req, csvContent)
	if err != nil {
		resp.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
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

// parseUploadRequest parses the form and CSV. When rows are invalid it returns
// one message per invalid row, numbered as in the file (the header is row 1).
func parseUploadRequest(r *http.Request) (dto.UploadTransactionRequest, []dto.UploadTransactionCSV, []string, error) {
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
		return dto.UploadTransactionRequest{}, nil, nil, errors.New("invalid instruction type")
	}

	if isInstructionStanding {
		instructionDateStr := r.FormValue("transaction_date")
		instructionDateInt, err := strconv.ParseInt(instructionDateStr, 10, 64)
		if err != nil {
			return dto.UploadTransactionRequest{}, nil, nil, err
		}
		instructionDate = time.Unix(instructionDateInt, 0)
	}

	totalRecordStr := r.FormValue("total_record")
	totalRecord, err := strconv.Atoi(totalRecordStr)
	if err != nil {
		return dto.UploadTransactionRequest{}, nil, nil, err
	}

	totalAmountStr := r.FormValue("total_amount")
	totalAmount, err := strconv.ParseFloat(totalAmountStr, 64)
	if err != nil {
		return dto.UploadTransactionRequest{}, nil, nil, err
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		return dto.UploadTransactionRequest{}, nil, nil, fmt.Errorf("fail to read form named file")
	}
	err = checkFileFormat(header, "csv")
	if err != nil {
		return dto.UploadTransactionRequest{}, nil, nil, err
	}
	csvContent, err = parseFileCSV(file, header)
	if err != nil {
		return dto.UploadTransactionRequest{}, nil, nil, err
	}

	if len(csvContent) < 2 {
		return dto.UploadTransactionRequest{}, nil, nil, errors.New("csv has no transaction rows")
	}

	var rowErrors []string
	for idx, row := range csvContent[1:] {
		rowNum := idx + 2 // 1-based, after the header row
		if len(row) != 4 {
			rowErrors = append(rowErrors, fmt.Sprintf("[row:%d] invalid column length", rowNum))
			continue
		}

		txRow := dto.UploadTransactionCSV{
			ToBankName:    strings.TrimSpace(row[0]),
			ToAccountNum:  strings.TrimSpace(row[1]),
			ToAccountName: strings.TrimSpace(row[2]),
		}
		switch {
		case txRow.ToBankName == "":
			rowErrors = append(rowErrors, fmt.Sprintf("[row:%d] invalid to_bank_name format", rowNum))
			continue
		case txRow.ToAccountNum == "":
			rowErrors = append(rowErrors, fmt.Sprintf("[row:%d] invalid to_account_no format", rowNum))
			continue
		case txRow.ToAccountName == "":
			rowErrors = append(rowErrors, fmt.Sprintf("[row:%d] invalid to_account_name format", rowNum))
			continue
		}

		transferAmount, err := strconv.ParseFloat(strings.TrimSpace(row[3]), 64)
		if err != nil || transferAmount <= 0 {
			rowErrors = append(rowErrors, fmt.Sprintf("[row:%d] invalid transfer_amount format", rowNum))
			continue
		}
		txRow.TransferAmount = transferAmount

		transactionData = append(transactionData, txRow)
	}

	if len(rowErrors) > 0 {
		return dto.UploadTransactionRequest{}, nil, rowErrors, errors.New("invalid csv content value")
	}

	return dto.UploadTransactionRequest{
		InstructionType: instructionType,
		TransactionDate: instructionDate,
		TotalAmount:     totalAmount,
		TotalRecord:     int32(totalRecord),
	}, transactionData, nil, nil
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
