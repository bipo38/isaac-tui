package isaac

import (
	"isaac-scrapper/config"
	"isaac-scrapper/internal/utils"

	"github.com/gocolly/colly"
)

type Transformation struct {
	name, id_game, effect, image string
	extension                    Extension
}

func CreateTransformationCsv() {

	var t Transformation

	writer, file := utils.CreateCsv(t, config.Transformation["csvRoute"], config.Transformation["csvName"])
	transformations := getTransformations()

	for _, v := range transformations {

		transformation := []string{
			v.name,
			v.id_game,
			v.effect,
			v.image,
			string(v.extension),
		}

		writer.Write(transformation)
	}

	defer file.Close()

	defer writer.Flush()

}

func getTransformations() []Transformation {

	collector := colly.NewCollector()

	var transformations []Transformation

	collector.OnHTML(config.Default["tableNode"], func(h *colly.HTMLElement) {

		transformation := newTransformation(h)

		transformations = append(transformations, transformation)
	})

	collector.Visit(config.Transformation["url"])

	return transformations

}

func newTransformation(el *colly.HTMLElement) Transformation {
	transformation := Transformation{
		name:      el.ChildAttr("td:nth-child(2)", "data-sort-value"),
		id_game:   el.ChildAttr("td:nth-child(1)", "data-sort-value"),
		effect:    el.ChildText("td:nth-child(4)>p"),
		image:     el.ChildAttr("td:nth-child(3)>a>img", "data-image-key"),
		extension: ParseExtension(el.ChildAttr("td:nth-child(2)>img", "title")),
	}

	imgUrl := el.ChildAttr("td:nth-child(3)>a>img", "data-src")
	utils.DownloadImage(imgUrl, config.Transformation["imgRoute"], transformation.image)
	// if err != nil {
	// 	fmt.Println("Failed to download image")
	// }

	return transformation

}
