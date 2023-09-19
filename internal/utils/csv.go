package utils

import (
	"encoding/csv"
	"os"
	"reflect"
)

func getHeaders(element interface{}) []string {

	structype := reflect.TypeOf(element)

	var headers []string

	for i := 0; i < structype.NumField(); i++ {

		headers = append(headers, structype.Field(i).Name)
	}

	return headers
}

func CreateCsv(category interface{}, fPath, fName string) (*csv.Writer, *os.File) {
	headers := getHeaders(category)

	route, fPath := routeParser(fPath, fName)
	createDirs(route)
	file := createFile(fPath)

	writer := csv.NewWriter(file)
	writer.Write(headers)

	return writer, file

}
