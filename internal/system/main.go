package system

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
)

func createDirs(path string) {

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("Error creating dirs %s: %s\n", path, err)
	}
}

func createFile(path string) *os.File {

	var file *os.File

	if _, err := os.Stat(path); err != nil {

		file, err = os.Create(path)
		if err != nil {
			log.Fatalf("Failed to create file: %s: %s\n ", path, err)
		}

	}

	return file
}

func getHeaders[C any](t C) []string {
	structype := reflect.TypeOf(t)

	var headers []string

	for i := 0; i < structype.NumField(); i++ {

		headers = append(headers, structype.Field(i).Name)
	}

	return headers
}

func CreateCsv[C any](category C, fPath, fName string) (*csv.Writer, *os.File) {
	defaultRoute := "isaac/"
	headers := getHeaders(category)
	route := fmt.Sprintf("%s%s/", defaultRoute, fPath)
	fPath = fmt.Sprintf("%s/%s", route, fName)

	createDirs(route)
	file := createFile(fPath)

	writer := csv.NewWriter(file)
	writer.Write(headers)

	return writer, file

}
