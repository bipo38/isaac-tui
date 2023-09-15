package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	// isaac.CreateTransformationCsv()
	// isaac.CreateTrinketsCsv()
	// isaac.CreateItemsCsv()
<<<<<<< Updated upstream
	isaac.CreateCharactersCsv()
=======
	// isaac.CreateCharactersCsv()
	// isaac.CreatePillsCsv()
	// isaac.CreateBossesCsv()

	collector := colly.NewCollector()

	collector.OnHTML("div.main-container>div.resizable-container>div.has-right-rail>main.page__main>div#content>div#mw-content-text>div.mw-parser-output", func(h *colly.HTMLElement) {
		fmt.Println(h.ChildAttr("img", "title"))
	})

	collector.Visit("https://bindingofisaacrebirth.fandom.com/wiki/All_Bosses_(Bosses)#Bosses")
>>>>>>> Stashed changes

}
