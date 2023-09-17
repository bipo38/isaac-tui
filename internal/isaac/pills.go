package isaac

import (
	"encoding/csv"
	"isaac-scrapper/internal/system"
	"strings"

	"github.com/gocolly/colly"
)

type Pill struct {
	name, effect, horse_effect, class, image string
	extension                                Extension
}

func CreatePillsCsv() {

	var t Pill

	fName := "pills.csv"
	route := defaultRoute + "pills/"
	fullRoute := route + fName

	pills := scrapingPills()
	headers := GetHeaders(t)
	system.CreateDirs(route)
	file := system.CreateFile(fullRoute)

	writer := csv.NewWriter(file)
	writer.Write(headers)

	for _, v := range pills {

		pill := []string{
			v.name,
			v.effect,
			v.horse_effect,
			v.class,
			v.image,
			string(v.extension),
		}

		writer.Write(pill)
	}

	defer file.Close()

	defer writer.Flush()

}

func scrapingPills() []Pill {

	collector := colly.NewCollector()

	var pills []Pill
	var extension string

	collector.OnHTML(TableNode, func(h *colly.HTMLElement) {
		pill := newPill(h, &extension)

		if pill.name == "" || strings.Contains(pill.name, "https") {
			return
		}

		pills = append(pills, pill)
	})

	collector.Visit(globaLink + PILLS)

	return pills
}

func newPill(el *colly.HTMLElement, extension *string) Pill {

	scrapingExtension := el.ChildAttr("th>b>a", "title")

	if scrapingExtension != "" {
		*extension = scrapingExtension
	}

	return Pill{
		name:         el.ChildText("td:nth-child(2)"),
		class:        el.ChildText("td:nth-child(3)"),
		effect:       el.ChildText("td:nth-child(4)"),
		horse_effect: el.ChildText("td:last-child"),
		image:        "image",
		extension:    ParseExtension(*extension),
	}
}
