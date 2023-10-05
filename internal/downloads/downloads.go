package downloads

import (
	"errors"
	"io"
	"isaac-scrapper/internal/manipulation"
	"isaac-scrapper/internal/utils"
	"net/http"
)

func Image(url, fPath, fName string) (string, error) {

	route, filePath := manipulation.RouteParser(fPath, fName)

	if err := utils.CreateDirs(route); err != nil {
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

	file, err := utils.CreateFile(filePath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
