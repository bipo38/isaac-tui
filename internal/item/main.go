package item

import (
	"encoding/csv"
	"isaac-scrapper/internal/system"
	"isaac-scrapper/internal/utils"
	"log"
	"reflect"

	"github.com/gocolly/colly"
)

type item struct {
	name, id_game, quote, effect, unlock, image, quality, pool string
	extension                                                  utils.Extension
}

// var node = "div.main-container>div.resizable-container>div.has-right-rail>main.page__main>div#content>div#mw-content-text>div.mw-parser-output"

func GetItemsCsv(fName, path string) {

	var t item

	fullRoute := path + fName

	headers := t.GetHeaders()
	content := TrinektScraping()

	system.CreateDirs(path)
	file := system.CreateFile(fullRoute)

	writer := csv.NewWriter(file)
	writer.Write(headers)

	for _, v := range content {

		item := []string{
			v.name,
			v.id_game,
			v.quote,
			v.effect,
			v.unlock,
			v.image,
			v.quality,
			v.pool,
			string(v.extension),
		}

		writer.Write(item)
	}

	defer file.Close()

	defer writer.Flush()

}

func (t item) GetHeaders() []string {

	structype := reflect.TypeOf(t)

	var headers []string

	for i := 0; i < structype.NumField(); i++ {

		headers = append(headers, structype.Field(i).Name)
	}

	return headers

}

func TrinektScraping() []item {

	collector := colly.NewCollector()

	node := "div.main-container>div.resizable-container>div.has-right-rail>main.page__main>div#content>div#mw-content-text>div.mw-parser-output>table.wikitable>tbody>tr.row-collectible"
	url := "https://bindingofisaacrebirth.fandom.com/wiki/Items"

	var items []item

	collector.OnHTML(node, func(el *colly.HTMLElement) {

		link := "https://bindingofisaacrebirth.fandom.com" + string(el.ChildAttr("td:nth-child(1)>a", "href"))

		item := item{
			name:      el.ChildAttr("td:nth-child(1)", "data-sort-value"),
			id_game:   el.ChildText("td:nth-child(2)"),
			quote:     el.ChildText("td:nth-child(4)"),
			effect:    el.ChildText("td:nth-child(5)"),
			image:     "imagenes3",
			quality:   el.ChildText("td:nth-child(6)"),
			extension: utils.ParseExtension(el.ChildAttr("td:nth-child(1)>img", "title")),
		}

		subcollector := colly.NewCollector()

		subcollector.OnHTML("div.main-container>div.resizable-container>div.has-right-rail>main.page__main>div#content>div#mw-content-text>div.mw-parser-output>aside", func(h *colly.HTMLElement) {
			unlock := h.ChildText("div[data-source=\"unlocked by\"]>div")

			pool := h.ChildText("div[data-source=\"alias\"]>div>div.item-pool-list")

			item.pool = pool

			if unlock != "" {
				item.unlock = unlock
			} else {
				item.unlock = "Unlocked"
			}

		})

		subcollector.Visit(link)

		items = append(items, item)

	})

	if err := collector.Visit(url); err != nil {
		log.Fatalf("Failed scraping the url %s: %s", url, err)
	}

	return items
}
