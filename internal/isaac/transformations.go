package isaac

import (
	"isaac-scrapper/internal/system"

	"github.com/gocolly/colly"
)

type Transformation struct {
	name, id_game, effect, image string
	extension                    Extension
}

func CreateTransformationCsv() {

	var t Transformation

	writer, file := system.CreateCsv(t, "transformations", "transformations.csv")
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

	collector.OnHTML(TableNode, func(h *colly.HTMLElement) {

		transformation := newTransformation(h)

		transformations = append(transformations, transformation)
	})

	collector.Visit(globaLink + TRANSFORMATIONS)

	return transformations

}

func newTransformation(el *colly.HTMLElement) Transformation {
	transformation := Transformation{
		name:      el.ChildAttr("td:nth-child(2)", "data-sort-value"),
		id_game:   el.ChildAttr("td:nth-child(1)", "data-sort-value"),
		effect:    el.ChildText("td:nth-child(4)>p"),
		image:     "imagenes",
		extension: ParseExtension(el.ChildAttr("td:nth-child(2)>img", "title")),
	}

	return transformation

}
