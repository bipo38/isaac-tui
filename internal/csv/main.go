package csv

import (
	"encoding/csv"
	"log"
	"os"
)

func newCsvFile(name string, headers []string) *csv.Writer {

	fName := name
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Failed to create file: %q: %s\n ", fName, err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(headers)

	return writer

}
