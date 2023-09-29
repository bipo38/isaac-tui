package main

import (
	"fmt"
	"isaac-scrapper/internal/isaac"
	"time"
)

func main() {
	// var wg sync.WaitGroup

	start := time.Now()

	// wg.Add(3)

	isaac.CreateTrinketsCsv()
	// isaac.CreatePillsCsv()
	// isaac.CreateCharactersCsv()

	// wg.Wait()

	duration := time.Since(start)

	fmt.Println(duration)
}
