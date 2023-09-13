package isaac2

import (
	"fmt"
	"isaac-scrapper/internal/utils"

	"github.com/gocolly/colly"
)

type Trinket struct {
	name, id_game, quote, effect, unlock, image string
	extension                                   utils.Extension
}

func GetTrinkets() []Trinket {

	collector := getScrapperInstance().collector

	// var trinkets []Trinket

	collector.OnHTML(mainNode, func(h *colly.HTMLElement) {

		// trinket := newTrinket()

		h.ForEach(mainNode, func(_ int, el *colly.HTMLElement) {
			fmt.Println(el.ChildText("div[data-source=\"unlocked by\"]>div"))

			h.Request.Visit("https://bindingofisaacrebirth.fandom.com/wiki/Butt_Penny")
		})

		// trinkets = append(trinkets, newTrinket(h))
	})

	collector.Visit(globaLink + "Trinkets")

	return trinkets

}

func newTrinket(h *colly.HTMLElement) Trinket {

	trinket := Trinket{}
	trinket.setName(h)

	return trinket
}

func (trinket *Trinket) setName(h *colly.HTMLElement) {
	trinket.name = "hola"
}

// func getPath(h *colly.HTMLElement) string {

// }
