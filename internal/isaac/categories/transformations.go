package categories

import (
	"log"

	"github.com/bipo38/isaac-tui/config"
	"github.com/bipo38/isaac-tui/internal/creates"
	"github.com/bipo38/isaac-tui/internal/downloads"
	"github.com/bipo38/isaac-tui/internal/isaac/parsers"
	"github.com/gocolly/colly"
)

type Transformation struct {
	name, id_game, effect, image, extension string
}

func CreateTransformationsCsv() error {

	transformations, err := scrapingTranformations()
	if err != nil {
		return err
	}

	csv := creates.Csv[Transformation]{
		Name:     config.Transformation["csvName"],
		Path:     config.Transformation["csvRoute"],
		Category: transformations,
	}

	if err := csv.Write(); err != nil {
		return err
	}

	return nil

}

func scrapingTranformations() ([]Transformation, error) {

	c := colly.NewCollector()

	var transformations []Transformation

	c.OnHTML(config.Default["tableNode"], func(el *colly.HTMLElement) {

		t, err := newTransformation(el)
		if err != nil {
			log.Printf("error creating transformation %v", err)
			return
		}

		transformations = append(transformations, *t)
	})

	if err := c.Visit(config.Transformation["url"]); err != nil {
		return nil, err
	}

	return transformations, nil

}

func newTransformation(el *colly.HTMLElement) (*Transformation, error) {

	t := Transformation{
		name:      el.ChildAttr("td:nth-child(2)", "data-sort-value"),
		id_game:   el.ChildAttr("td:nth-child(1)", "data-sort-value"),
		effect:    el.ChildText("td:nth-child(4)>p"),
		extension: parsers.Extension(el.ChildAttr("td:nth-child(2)>img", "title")),
	}

	if err := setTransformationImage(el, &t); err != nil {
		return nil, err
	}

	return &t, nil

}

func setTransformationImage(el *colly.HTMLElement, transformation *Transformation) error {

	url := el.ChildAttr("td:nth-child(3)>a>img", "data-src")
	name := el.ChildAttr("td:nth-child(3)>a>img", "data-image-key")

	p, err := downloads.Image(url, config.Transformation["imgFolder"], name)
	if err != nil {
		return err
	}

	transformation.image = p

	return nil

}
