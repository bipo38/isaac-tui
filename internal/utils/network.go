package utils

import (
	"errors"
	"io"
	"net/http"
)

func DownloadImage(url, fPath, fName string) error {

	route, filePath := RouteParser(fPath, fName)

	CreateDirs(route)

	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}

	file := CreateFile(filePath)

	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
