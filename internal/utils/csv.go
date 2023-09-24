package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func CreateCsv(category interface{}, fPath, fName string) (*csv.Writer, *os.File) {
	headers := GetHeaders(category)

	fName = ParserFileName(fName, "csv")

	route, fPath := RouteParser(fPath, fName)
	CreateDirs(route)
	file := CreateFile(fPath)

	writer := csv.NewWriter(file)
	if err := writer.Write(headers); err != nil {
		log.Fatalf("Falied creating csv file: %s , Path: %s", err, fPath)
	}

	return writer, file

}
