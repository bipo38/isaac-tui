package categories

import (
	"errors"
	"isaac-scrapper/config"
	"isaac-scrapper/internal/creates"
	"isaac-scrapper/internal/isaac/parsers"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Pill struct {
	name, effect, horse_effect, class, image, extension string
}

func CreatePillsCsv() error {

	pills, err := scrapingPills()
	if err != nil {
		return err
	}

	csv := creates.Csv[Pill]{
		Name:     config.Pill["csvName"],
		Path:     config.Pill["csvRoute"],
		Category: pills,
	}

	if err := csv.Write(); err != nil {
		return err
	}

	return nil

}

func scrapingPills() ([]Pill, error) {

	c := colly.NewCollector()

	var pills []Pill
	var ext string

	c.OnHTML(config.Default["tableNode"], func(h *colly.HTMLElement) {
		pill, err := newPill(h, &ext)
		if err != nil {
			log.Printf("error creating pill: %v", err)
			return
		}

		pills = append(pills, *pill)
	})

	if err := c.Visit(config.Pill["url"]); err != nil {
		return nil, err
	}

	return pills, nil
}

func newPill(el *colly.HTMLElement, ext *string) (*Pill, error) {

	scrapingExtension := el.ChildAttr("th>b>a", "title")

	if scrapingExtension != "" {
		*ext = scrapingExtension
	}

	p := Pill{
		name: el.ChildText("td:nth-child(2)"),
	}
	if p.name == "" || strings.Contains(p.name, "https") {
		return nil, errors.New("name is empty")
	}

	p.class = el.ChildText("td:nth-child(3)")
	p.effect = el.ChildText("td:nth-child(4)")
	p.horse_effect = el.ChildText("td:last-child")
	p.image = "image"
	p.extension = parsers.Extension(*ext)

	return &p, nil
}
