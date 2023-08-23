package transformation

import (
	"encoding/csv"
	"isaac-scrapper/internal/system"
	"isaac-scrapper/internal/utils"
	"log"
	"reflect"

	"github.com/gocolly/colly"
)

type transformation struct {
	name, id_game, effect, image string
	extension                    utils.Extension
}

func GetTransformationCsv(fName, path string) {

	fullRoute := path + fName

	headers := GetHeaders()
	content := TransformationScraping()

	system.CreateDirs(path)
	file := system.CreateFile(fullRoute)

	writer := csv.NewWriter(file)
	writer.Write(headers)

	for _, v := range content {

		transformation := []string{
			v.id_game,
			v.name,
			v.effect,
			v.image,
			string(v.extension),
		}

		writer.Write(transformation)
	}

	defer file.Close()

	defer writer.Flush()

}

func GetHeaders() []string {

	var transformation transformation

	structype := reflect.TypeOf(transformation)

	var headers []string

	for i := 0; i < structype.NumField(); i++ {

		headers = append(headers, structype.Field(i).Name)
	}

	return headers

}

func TransformationScraping() []transformation {

	collector := colly.NewCollector()

	node := "div.main-container>div.resizable-container>div.has-right-rail>main.page__main>div#content>div#mw-content-text>div.mw-parser-output>table.wikitable>tbody>tr.row-transformation"
	url := "https://bindingofisaacrebirth.fandom.com/wiki/Transformations"

	var transformations []transformation

	collector.OnHTML(node, func(el *colly.HTMLElement) {

		transformation := transformation{
			name:      el.ChildAttr("td:nth-child(1)", "data-sort-value"),
			id_game:   el.ChildAttr("td:nth-child(2)", "data-sort-value"),
			effect:    el.ChildText("td:nth-child(4)>p"),
			image:     "imagenes",
			extension: utils.ParseExtension(el.ChildAttr("td:nth-child(2)>img", "title")),
		}

		transformations = append(transformations, transformation)

	})

	if err := collector.Visit(url); err != nil {
		log.Fatalf("Failed scraping the url %s: %s", url, err)
	}

	return transformations
}
