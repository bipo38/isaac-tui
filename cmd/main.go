package main

import "isaac-scrapper/internal/transformation"

func main() {

	// collector := colly.NewCollector()

	// transformations := []string{}

	// collector.OnHTML("div.main-container>div.resizable-container>div.has-right-rail>main.page__main>div#content>div#mw-content-text>div.mw-parser-output>table.wikitable>tbody", func(e *colly.HTMLElement) {

	// 	transformations = e.ChildAttrs("tr>td:nth-child(2)>a", "href")

	// })

	// collector.Visit("https://bindingofisaacrebirth.fandom.com/wiki/Transformations")
	// scrap := transformation.TransformationScraping()
	// fmt.Println(len(scrap))
	// for _, v := range scrap {

	// 	fmt.Println(v)
	// 	fmt.Println()
	// }

	transformation.GetTransformationCsv()

}
