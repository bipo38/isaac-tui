package utils

import (
	"encoding/csv"
	"log"
	"os"
	"reflect"
	"strings"
)

func CreateCsv(category interface{}, fPath, fName string) (*csv.Writer, *os.File) {
	headers := getHeaders(category)

	route, fPath := routeParser(fPath, (fName + ".csv"))
	createDirs(route)
	file := createFile(fPath)

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

func parseCsvFileName(fName *string) string {
	extension := ".csv"

	if strings.Contains(*fName, extension) {
		return *fName
	}

	return *fName + extension
}
