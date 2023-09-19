package system

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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

	route, fPath := filesStarter(fPath, fName)
	createDirs(route)
	file := createFile(fPath)

	writer := csv.NewWriter(file)
	writer.Write(headers)

	return writer, file

}

func filesStarter(fRoute, fName string) (string, string) {
	defaultRoute := "isaac"

	route := fmt.Sprintf("%s/%s/", defaultRoute, fRoute)
	fPath := fmt.Sprintf("%s%s", route, fName)

	fmt.Println(fPath)

	return route, fPath

}

func DownloadFile(url, fPath, fName string) error {

	route, filePath := filesStarter(fPath, fName)

	createDirs(route)

	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	file := createFile(filePath)

	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
