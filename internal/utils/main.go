package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

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
