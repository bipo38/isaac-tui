package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	collector := colly.NewCollector()

	transformations := []string{}

	collector.OnHTML("div.main-container>div.resizable-container>div.has-right-rail>main.page__main>div#content>div#mw-content-text>div.mw-parser-output>table.wikitable>tbody", func(e *colly.HTMLElement) {

		transformations = e.ChildAttrs("tr>td:nth-child(2)>a", "href")

	})

	collector.Visit("https://bindingofisaacrebirth.fandom.com/wiki/Transformations")

	for _, v := range transformations {

		fmt.Println(v)
		fmt.Println()
	}

}
