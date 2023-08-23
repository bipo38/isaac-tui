package main

import (
	"isaac-scrapper/internal/item"
)

func main() {

	// transformation.GetTransformationCsv("transformation.csv", "isaac/transformations")
	// trinket.GetTrinektsCsv("trinkets.csv", "isaac/trinkets/")
	item.GetItemsCsv("items.csv", "isaac/items/")
}
