package isaac

import (
	"encoding/csv"
	"isaac-scrapper/internal/system"

	"github.com/gocolly/colly"
)

type Trinket struct {
	name, id_game, quote, effect, unlock, image string
	extension                                   Extension
}

func CreateTrinketsCsv() {

	var t Trinket

	fName := "trinkets.csv"
	route := defaultRoute + "trinkets/"
	fullRoute := route + fName

	trinkets := getTrinkets()
	headers := GetHeaders(t)
	system.CreateDirs(route)
	file := system.CreateFile(fullRoute)

	writer := csv.NewWriter(file)
	writer.Write(headers)

	for _, v := range trinkets {

		trinket := []string{
			v.name,
			v.id_game,
			v.quote,
			v.effect,
			v.unlock,
			v.image,
			string(v.extension),
		}

		writer.Write(trinket)
	}

	defer file.Close()

	defer writer.Flush()

}

func getTrinkets() []Trinket {

	collector := colly.NewCollector()

	var trinkets []Trinket

	collector.OnHTML(TableNode, func(h *colly.HTMLElement) {

		trinket := newTrinket(h.ChildAttr("a", "href"), h)

		trinkets = append(trinkets, trinket)
	})

	collector.Visit(globaLink + string(TRINKETS))

	return trinkets

}

func newTrinket(path string, el *colly.HTMLElement) Trinket {
	trinket := Trinket{
		name:    el.ChildAttr("td:nth-child(1)", "data-sort-value"),
		id_game: el.ChildText("td:nth-child(2)"),
		quote:   el.ChildText("td:nth-child(4)"),
		effect:  el.ChildText("td:nth-child(5)"),
		image:   "image",
	}

	collector := colly.NewCollector()

	collector.OnHTML(mainNode, func(h *colly.HTMLElement) {
		setTrinket(h, &trinket)
		setExtension(h, &trinket)

	})

	collector.Visit(globaLink + path)

	return trinket

}

func setTrinket(h *colly.HTMLElement, trinket *Trinket) {
	unlock := h.ChildText("div[data-source=\"unlocked by\"]>div")

	if unlock != "" {
		trinket.unlock = unlock
	} else {
		trinket.unlock = "Unlocked"
	}

}

func setExtension(h *colly.HTMLElement, trinket *Trinket) {
	extension := h.ChildAttr("div#context-page.context-box>img", "title")

	trinket.extension = ParseExtension(extension)
}
