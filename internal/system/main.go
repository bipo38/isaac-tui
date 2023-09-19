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

func routeParser(fRoute, fName string) (string, string) {
	defaultRoute := "isaac"

	route := fmt.Sprintf("%s/%s/", defaultRoute, fRoute)
	fPath := fmt.Sprintf("%s%s", route, fName)

	fmt.Println(fPath)

	return route, fPath

}
func DownloadImage(url, fPath, fName string) error {

	route, filePath := routeParser(fPath, fName)

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

func createDirs(path string) {
	exist, err := exist(path)
	if err != nil {
		log.Fatalf("Failed verify the existence of the folder: %s: %s\n ", path, err)
	}

	if exist {
		os.Remove(path)
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("Error creating dirs %s: %s\n", path, err)
	}

	return
}

func createFile(path string) *os.File {

	exist, err := exist(path)
	if err != nil {
		log.Fatalf("Failed verify the existence of the file: %s: %s\n ", path, err)
	}

	if exist {
		os.Remove(path)
	}

	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("Error creating file: %s: %s\n ", path, err)
	}

	return file
}

func exist(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
