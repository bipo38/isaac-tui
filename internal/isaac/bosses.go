package isaac

import (
	"encoding/csv"
	"isaac-scrapper/internal/system"

	"github.com/gocolly/colly"
)

type Boss struct {
	name, id_game, unlock, image string
	extension                    Extension
}

func CreateBossesCsv() {

	var t Boss

	fName := "bosses.csv"
	route := defaultRoute + "boss/"
	fullRoute := route + fName

	bosses := scrapingBosses()
	headers := GetHeaders(t)
	system.CreateDirs(route)
	file := system.CreateFile(fullRoute)

	writer := csv.NewWriter(file)
	writer.Write(headers)

	for _, v := range bosses {

		boss := []string{
			v.name,
			v.id_game,
			v.unlock,
			v.image,
			string(v.extension),
		}

		writer.Write(boss)
	}

	defer file.Close()

	defer writer.Flush()

}

func scrapingBosses() []Boss {

	collector := colly.NewCollector()

	var bosss []Boss

	collector.OnHTML(TableNode, func(h *colly.HTMLElement) {

		boss := newboss(h.ChildAttr("a", "href"), h)

		bosss = append(bosss, boss)
	})

	collector.Visit(globaLink + BOSSES)

	return bosss

}

func newboss(path string, el *colly.HTMLElement) Boss {
	boss := Boss{
		name:  el.ChildAttr("td>a", "title"),
		image: "image",
	}

	collector := colly.NewCollector()

	collector.OnHTML(mainNode, func(h *colly.HTMLElement) {
		setBossUnlock(h, &boss)
		setBossExtension(h, &boss)
		setBossId(h, &boss)

	})

	collector.Visit(globaLink + path)

	return boss

}

func setBossId(h *colly.HTMLElement, boss *Boss) {
	boss.id_game = h.ChildText("div[data-source=\"id\"]>div")
}

func setBossUnlock(h *colly.HTMLElement, boss *Boss) {
	unlock := h.ChildText("div[data-source=\"unlocked by\"]>div")

	if unlock != "" {
		boss.unlock = unlock
	} else {
		boss.unlock = "Unlocked"
	}

}

func setBossExtension(h *colly.HTMLElement, boss *Boss) {
	extension := h.ChildAttr("div#context-page.context-box>img", "title")

	boss.extension = ParseExtension(extension)
}
