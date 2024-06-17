package module

import (
	"encoding/csv"
	"errors"
	"io"
	"mime/multipart"
	"strings"
)

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
