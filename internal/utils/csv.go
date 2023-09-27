package utils

import (
	"encoding/csv"
	"os"
)

func CreateCsv(category interface{}, fPath, fName string) (*csv.Writer, *os.File, error) {
	headers := GetHeaders(category)

	fName = ParserFileName(fName, "csv")

	route, fPath := RouteParser(fPath, fName)

	if err := CreateDirs(route); err != nil {
		return nil, nil, err
	}

	file, err := CreateFile(fPath)
	if err != nil {
		return nil, nil, err
	}

	writer := csv.NewWriter(file)
	if err := writer.Write(headers); err != nil {
		return nil, nil, err
	}

	return writer, file, nil

}
