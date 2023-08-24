package main

import (
	"isaac-scrapper/internal/item"
)

func main() {

	// transformation.GetTransformationCsv("transformation.csv", "isaac/transformations")
	// trinket.GetTrinektsCsv("trinkets.csv", "isaac/trinkets/")
	item.GetItemsCsv("items.csv", "isaac/items/")

	// c := colly.NewCollector()

	// url := "div.main-container>div.resizable-container>div.has-right-rail>main.page__main>div#content>div#mw-content-text>div.mw-parser-output"

	// c.OnHTML(url, func(h *colly.HTMLElement) {

	// 	fmt.Println(h.ChildText("p:nth-child(4)>a"))
	// })

	// c.Visit("https://bindingofisaacrebirth.fandom.com/wiki/Abaddon")
}
