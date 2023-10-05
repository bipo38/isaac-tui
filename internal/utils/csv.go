package utils

import (
	"encoding/csv"
	"isaac-scrapper/internal/creates"
	"isaac-scrapper/internal/manipulation"
	"os"
)

func CreateCsv(category interface{}, fPath, fName string) (*csv.Writer, *os.File, error) {
	headers := manipulation.GetHeaders(category)

	fName = manipulation.ParserFileName(fName, "csv")

	_, fPath = manipulation.RouteParser(fPath, fName)

	file, err := creates.File(fPath)
	if err != nil {
		return nil, nil, err
	}

	writer := csv.NewWriter(file)
	if err := writer.Write(headers); err != nil {
		return nil, nil, err
	}

	return writer, file, nil

}
