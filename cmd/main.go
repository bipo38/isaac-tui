package main

import (
	"errors"
	"fmt"
	"isaac-scrapper/internal/isaac"
)

func main() {

	// isaac.CreateTransformationCsv()
	// isaac.CreateTrinketsCsv()
	// isaac.CreateItemsCsv()
	if err := isaac.CreateCharactersCsv(); err != nil {
		unwrapedError := errors.Unwrap(err)
		fmt.Errorf("error creating file: %w", unwrapedError)
	}
	// isaac.CreatePillsCsv()

}
