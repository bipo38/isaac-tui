package isaac

import (
	"fmt"
	"isaac-scrapper/config"
	"isaac-scrapper/internal/utils"

	"github.com/gocolly/colly"
)

type Trinket struct {
	name, id_game, quote, effect, unlock, image, extension string
}

func CreateTrinketsCsv() {

	var t Trinket

	writer, file := utils.CreateCsv(t, "trinkets", "trinkets.csv")
	trinkets := getTrinkets()

	for _, v := range trinkets {

		trinket := []string{
			v.name,
			v.id_game,
			v.quote,
			v.effect,
			v.unlock,
			v.image,
			v.extension,
		}

		writer.Write(trinket)
	}

	defer file.Close()

	defer writer.Flush()

}

func getTrinkets() []Trinket {

	collector := colly.NewCollector()

	var trinkets []Trinket

	collector.OnHTML(config.Default["tableNode"], func(el *colly.HTMLElement) {

		trinket := newTrinket(el)

		trinkets = append(trinkets, trinket)
	})

	collector.Visit(config.Trinket["url"])

	return trinkets

}

func newTrinket(el *colly.HTMLElement) Trinket {
	urlPath := el.ChildAttr("a", "href")

	trinket := Trinket{
		name:    el.ChildAttr("td:nth-child(1)", "data-sort-value"),
		id_game: el.ChildText("td:nth-child(2)"),
		quote:   el.ChildText("td:nth-child(4)"),
		effect:  el.ChildText("td:nth-child(5)"),
		image:   "image",
	}

	collector := colly.NewCollector()

	collector.OnHTML(config.Default["mainNode"], func(h *colly.HTMLElement) {
		setTrinketUnlock(h, &trinket)
		setTrinketExtension(h, &trinket)

	})

	collector.Visit(fmt.Sprintf("%s%s", config.Default["url"], urlPath))

	return trinket

}

func setTrinketUnlock(h *colly.HTMLElement, trinket *Trinket) {
	unlock := h.ChildText("div[data-source=\"unlocked by\"]>div")

	if unlock != "" {
		trinket.unlock = unlock
	} else {
		trinket.unlock = "Unlocked"
	}

}

func setTrinketExtension(h *colly.HTMLElement, trinket *Trinket) {
	extension := h.ChildAttr("div#context-page.context-box>img", "title")
	trinket.extension = ParseExtension(extension)
}
