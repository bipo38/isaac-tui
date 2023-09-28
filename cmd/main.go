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
	// if err := isaac.CreateCharactersCsv(); err != nil {
	// 	unwrapedError := errors.Unwrap(err)
	// 	fmt.Printf("error creating file: %v", unwrapedError)

	// }

	if err := isaac.CreateTransformationCsv(); err != nil {
		unwrapedError := errors.Unwrap(err)
		fmt.Printf("error creating file: %v", unwrapedError)
	}
	// isaac.CreatePillsCsv()

}
