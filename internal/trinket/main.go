package trinket

import (
	"encoding/csv"
	"isaac-scrapper/internal/system"
	"isaac-scrapper/internal/utils"
	"log"
	"reflect"

	"github.com/gocolly/colly"
)

type trinket struct {
	name, id_game, quote, effect, unlock, image string
	extension                                   utils.Extension
}

func GetTrinektsCsv(fName, path string) {

	var t trinket

	fullRoute := path + fName

	headers := t.GetHeaders()
	content := TrinektScraping()

	system.CreateDirs(path)
	file := system.CreateFile(fullRoute)

	writer := csv.NewWriter(file)
	writer.Write(headers)

	for _, v := range content {

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

func (t trinket) GetHeaders() []string {

	structype := reflect.TypeOf(t)

	var headers []string

	for i := 0; i < structype.NumField(); i++ {

		headers = append(headers, structype.Field(i).Name)
	}

	return headers

}

func TrinektScraping() []trinket {

	collector := colly.NewCollector()

	node := "div.main-container>div.resizable-container>div.has-right-rail>main.page__main>div#content>div#mw-content-text>div.mw-parser-output>table.wikitable>tbody>tr.row-trinket"
	url := "https://bindingofisaacrebirth.fandom.com/wiki/Trinkets"

	var trinkets []trinket

	collector.OnHTML(node, func(el *colly.HTMLElement) {

		link := "https://bindingofisaacrebirth.fandom.com/" + string(el.ChildAttr("td:nth-child(1)>a", "href"))

		trinket := trinket{
			name:      el.ChildAttr("td:nth-child(2)", "data-sort-value"),
			id_game:   el.ChildText("td:nth-child(2)"),
			quote:     el.ChildText("td:nth-child(4)"),
			effect:    el.ChildText("td:nth-child(5)"),
			image:     "imagenes3",
			extension: utils.ParseExtension(el.ChildAttr("td:nth-child(1)>img", "title")),
		}

		subcollector := colly.NewCollector()

		subcollector.OnHTML("div.main-container>div.resizable-container>div.has-right-rail>main.page__main>div#content>div#mw-content-text>div.mw-parser-output>aside", func(h *colly.HTMLElement) {
			unlock := h.ChildText("div[data-source=\"unlocked by\"]>div")

			if unlock != "" {
				trinket.unlock = unlock
			} else {
				trinket.unlock = "Unlocked"
			}

		})

		subcollector.Visit(link)

		trinkets = append(trinkets, trinket)

	})

	if err := collector.Visit(url); err != nil {
		log.Fatalf("Failed scraping the url %s: %s", url, err)
	}

	return trinkets
}
