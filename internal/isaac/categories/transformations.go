package categories

import (
	"isaac-scrapper/config"
	"isaac-scrapper/internal/creates"
	"isaac-scrapper/internal/downloads"
	"isaac-scrapper/internal/isaac/parsers"
	"log"

	"github.com/gocolly/colly"
)

type Transformation struct {
	name, id_game, effect, image, extension string
}

func CreateTransformationsCsv() error {

	var t Transformation

	writer, file, err := creates.Csv(t, config.Transformation["csvRoute"], config.Transformation["csvName"])
	if err != nil {
		return err
	}

	transformations, err := scrapingTranformations()
	if err != nil {
		return err
	}

	for _, v := range transformations {

		transformation := []string{
			v.name,
			v.id_game,
			v.effect,
			v.image,
			v.extension,
		}

		writer.Write(transformation)
	}

	defer file.Close()

	defer writer.Flush()

	return nil
}

func scrapingTranformations() ([]Transformation, error) {

	collector := colly.NewCollector()

	var transformations []Transformation

	collector.OnHTML(config.Default["tableNode"], func(el *colly.HTMLElement) {

		transformation, err := newTransformation(el)
		if err != nil {
			log.Printf("error creating transformation %v", err)
			return
		}

		transformations = append(transformations, *transformation)
	})

	if err := collector.Visit(config.Transformation["url"]); err != nil {
		return nil, err
	}

	return transformations, nil

}

func newTransformation(el *colly.HTMLElement) (*Transformation, error) {

	transformation := Transformation{
		name:      el.ChildAttr("td:nth-child(2)", "data-sort-value"),
		id_game:   el.ChildAttr("td:nth-child(1)", "data-sort-value"),
		effect:    el.ChildText("td:nth-child(4)>p"),
		extension: parsers.ParseExtension(el.ChildAttr("td:nth-child(2)>img", "title")),
	}

	if err := setTransformationImage(el, &transformation); err != nil {
		return nil, err
	}

	return &transformation, nil

}

func setTransformationImage(el *colly.HTMLElement, transformation *Transformation) error {

	imgUrl := el.ChildAttr("td:nth-child(3)>a>img", "data-src")
	imgName := el.ChildAttr("td:nth-child(3)>a>img", "data-image-key")

	imgPath, err := downloads.Image(imgUrl, config.Transformation["imgFolder"], imgName)
	if err != nil {
		return err
	}

	transformation.image = imgPath

	return nil

}
