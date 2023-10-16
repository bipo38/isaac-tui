package downloads

import (
	"errors"
	"io"
	"net/http"

	"github.com/bipo38/isaac-tui/internal/creates"
	"github.com/bipo38/isaac-tui/internal/manipulation"
)

func Image(url, fp, fn string) (string, error) {

	_, fp = manipulation.RouteParser(fp, fn)

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.New("received non 200 response code")
	}

	f, err := creates.File(fp)
	if err != nil {
		return "", err
	}

	defer f.Close()

	_, err = io.Copy(f, res.Body)
	if err != nil {
		return "", err
	}

	return fp, nil
}
