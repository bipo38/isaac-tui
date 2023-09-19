package utils

import (
	"log"
	"os"
)

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
