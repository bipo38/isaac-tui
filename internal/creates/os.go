package creates

import (
	"os"
	"strings"
)

func Dirs(path string) error {
	exist, err := ExistPath(path)
	if err != nil {
		return err
	}

	if exist {
		return nil
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {

		return err
	}

	return nil

}

func File(fPath string) (*os.File, error) {

	splitPath := strings.Split(fPath, "/")

	fileRoute := strings.Join(splitPath[0:len(splitPath)-1], "/")

	exist, err := ExistPath(fPath)
	if err != nil {
		return nil, err
	}

	if exist {
		os.Remove(fPath)
	} else {
		Dirs(fileRoute)
	}

	file, err := os.Create(fPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func ExistPath(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
