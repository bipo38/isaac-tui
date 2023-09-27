package utils

import (
	"os"
)

func CreateDirs(path string) error {
	exist, err := ExistPath(path)
	if err != nil {
		return err
	}

	if exist {
		os.Remove(path)
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {

		return err
	}

	return nil

}

func CreateFile(path string) (*os.File, error) {

	exist, err := ExistPath(path)
	if err != nil {
		return nil, err
	}

	if exist {
		os.Remove(path)
	}

	file, err := os.Create(path)
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
