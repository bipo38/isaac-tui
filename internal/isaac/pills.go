package isaac

import (
	"isaac-scrapper/config"
	"isaac-scrapper/internal/utils"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Pill struct {
	name, effect, horse_effect, class, image, extension string
}

func CreatePillsCsv() error {

	var t Pill

	writer, file, err := utils.CreateCsv(t, config.Pill["csvRoute"], config.Pill["csvName"])
	if err != nil {
		return err
	}

	pills, err := scrapingPills()
	if err != nil {
		return err
	}

	for _, v := range pills {

		pill := []string{
			v.name,
			v.effect,
			v.horse_effect,
			v.class,
			v.image,
			string(v.extension),
		}

		if err := writer.Write(pill); err != nil {
			continue
		}

	}

	defer file.Close()

	defer writer.Flush()

	return nil

}

func scrapingPills() ([]Pill, error) {

	collector := colly.NewCollector()

	var pills []Pill
	var extension string

	collector.OnHTML(config.Default["tableNode"], func(h *colly.HTMLElement) {
		pill := newPill(h, &extension)

		if pill.name == "" || strings.Contains(pill.name, "https") {
			log.Println("skipping pill")
			return
		}

		pills = append(pills, pill)
	})

	if err := collector.Visit(config.Pill["url"]); err != nil {
		return nil, err
	}

	return pills, nil
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
		extension:    parseExtension(*extension),
	}
}
