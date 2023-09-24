package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

func CreateCsv(category interface{}, fPath, fName string) (*csv.Writer, *os.File) {
	headers := getHeaders(category)

	fName = parserCsvFileName(fName)

	route, fPath := RouteParser(fPath, fName)
	CreateDirs(route)
	file := CreateFile(fPath)

	writer := csv.NewWriter(file)
	if err := writer.Write(headers); err != nil {
		log.Fatalf("Falied creating csv file: %s , Path: %s", err, fPath)
	}

	return writer, file

}

func getHeaders(element interface{}) []string {

	structype := reflect.TypeOf(element)

	var headers []string

	for i := 0; i < structype.NumField(); i++ {

		headers = append(headers, structype.Field(i).Name)
	}

	return headers
}

func parserCsvFileName(fName string) string {
	extension := ".csv"

	if strings.HasSuffix(fName, extension) {
		return fName
	}

	return fmt.Sprintf("%s%s", fName, extension)
}
