package main

import (
	"fmt"
	"isaac-scrapper/internal/isaac2"
)

func main() {

	fmt.Println(isaac2.GetTrinkets())

	// transformation.GetTransformationCsv("transformation.csv", "isaac/transformations")
	// trinket.GetTrinektsCsv("trinkets.csv", "isaac/trinkets/")
	// item.GetItemsCsv("items.csv", "isaac/items/")

	// // var cards []string

	// c := colly.NewCollector()
	// fmt.Println("hola")

	// c.OnHTML("div.main-container>div.resizable-container>div.has-right-rail>main.page__main>div#content>div#mw-content-text>div.mw-parser-output>table.wikitable>tbody>tr", func(h *colly.HTMLElement) {
	// 	// fmt.Println(h.ChildAttrs("a", "title"))

	// 	fmt.Println(h.ChildAttrs("a", "title"))

	// 	// if name == "" {
	// 	// 	return
	// 	// }

	// 	// cards = append(cards, name)
	// })

	// print(cards)

	// c.Visit("https://bindingofisaacrebirth.fandom.com/wiki/Cards_and_Runes")

	// isaac.StartScraping(isaac.ITEMS)

}

// ITEMS           Category = "/Items"
// TRINKETS                 = "/Trinkets"
// TRANSFORMATIONS          = "/Transformations"
// BOSSES                   = "/All_Bosses_(Bosses)"
// CHARACTERS               = "/Characters"
// CARDS                    = "/Cards_and_Runes"
// PILLS                    = "/Pills"
