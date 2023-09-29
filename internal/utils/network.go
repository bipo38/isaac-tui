package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func DownloadImage(url, fPath, fName string) (string, error) {

	route, filePath := RouteParser(fPath, fName)

	if err := CreateDirs(route); err != nil {
		return "", err
	}

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", errors.New("received non 200 response code")
	}

	file, err := CreateFile(filePath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	imgPath := fmt.Sprintf("./%s/%s", fPath, fName)

	return imgPath, nil
}
