package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	// transformation.GetTransformationCsv("transformation.csv", "isaac/transformations")
	// trinket.GetTrinektsCsv("trinkets.csv", "isaac/trinkets/")
	// item.GetItemsCsv("items.csv", "isaac/items/")

	// isaac.DoScraping(isaac.BOSSES, isaac.BOSS_PAGE)

	c := colly.NewCollector()
	fmt.Println("hola")

	url := "div.main-container>div.resizable-container>div.has-right-rail>main.page__main>div#content>div#mw-content-text>div.mw-parser-output>div.table-wide"

	c.OnHTML(url, func(h *colly.HTMLElement) {

		fmt.Println(h.Text)
	})

	c.Visit("https://bindingofisaacrebirth.fandom.com/wiki/All_Bosses_(Bosses)")
}
