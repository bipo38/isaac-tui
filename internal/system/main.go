package system

import (
	"log"
	"os"
)

func CreateDirs(path string) {

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("Error creating dirs %s: %s\n", path, err)
	}
}

func CreateFile(path string) *os.File {

	var file *os.File

	if _, err := os.Stat(path); err != nil {

		file, err = os.Create(path)
		if err != nil {
			log.Fatalf("Failed to create file: %s: %s\n ", path, err)
		}

	}

	return file
}
