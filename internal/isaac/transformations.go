package isaac

import (
	"isaac-scrapper/config"
	"isaac-scrapper/internal/utils"
	"log"

	"github.com/gocolly/colly"
)

type Transformation struct {
	name, id_game, effect, image, extension string
}

func CreateTransformationCsv() error {

	var t Transformation

	writer, file, err := utils.CreateCsv(t, config.Transformation["csvRoute"], config.Transformation["csvName"])
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

	collector.OnHTML(config.Default["tableNode"], func(h *colly.HTMLElement) {

		transformation, err := newTransformation(h)
		if err != nil {
			log.Printf("error creating transformation: %v", err)
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
		image:     el.ChildAttr("td:nth-child(3)>a>img", "data-image-key"),
		extension: parseExtension(el.ChildAttr("td:nth-child(2)>img", "title")),
	}

	imgUrl := el.ChildAttr("td:nth-child(3)>a>img", "data-src")
	if err := utils.DownloadImage(imgUrl, config.Transformation["imgRoute"], transformation.image); err != nil {
		return nil, err
	}
	return &transformation, nil

}
