package isaac

import (
	"encoding/csv"
	"isaac-scrapper/internal/system"

	"github.com/gocolly/colly"
)

type Transformation struct {
	name, id_game, effect, image string
	extension                    Extension
}

func CreateTransformationCsv() {

	var t Transformation

	fName := "transformations.csv"
	route := defaultRoute + "transformations/"
	fullRoute := route + fName

	transformations := getTransformations()
	headers := GetHeaders(t)
	system.CreateDirs(route)
	file := system.CreateFile(fullRoute)

	writer := csv.NewWriter(file)
	writer.Write(headers)

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

		transformation := newTransformation(h.ChildAttr("a", "href"), h)

		transformations = append(transformations, transformation)
	})

	collector.Visit(globaLink + string(TRANSFORMATIONS))

	return transformations

}

func newTransformation(path string, el *colly.HTMLElement) Transformation {
	transformation := Transformation{
		name:      el.ChildAttr("td:nth-child(2)", "data-sort-value"),
		id_game:   el.ChildAttr("td:nth-child(1)", "data-sort-value"),
		effect:    el.ChildText("td:nth-child(4)>p"),
		image:     "imagenes",
		extension: ParseExtension(el.ChildAttr("td:nth-child(2)>img", "title")),
	}

	return transformation

}
