package transformation

import (
	"encoding/csv"
	"isaac-scrapper/internal/utils"
	"log"
	"os"
	"reflect"

	"github.com/gocolly/colly"
)

type transformation struct {
	name, id_game, effect, image string
	extension                    utils.Extension
}

func GetTransformationCsv() {

	headers := GetHeaders()
	content := TransformationScraping()
	file := "transformation.csv"

	NewCsvFile(file, "isaac-data/transformations/", headers, content)

}

func NewCsvFile(name, path string, headers []string, content []transformation) {

	var file *os.File

	fPath := path
	fName := name
	fullRoute := fPath + fName

	if err := os.MkdirAll(fPath, os.ModePerm); err != nil {
		log.Fatal("Failed to create dirs")
	}

	if _, err := os.Stat(fullRoute); err != nil {

		file, err = os.Create(fullRoute)
		if err != nil {
			log.Fatalf("Failed to create file: %q: %s\n ", fName, err)
		}

		defer file.Close()
	}

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

	defer writer.Flush()

}

func GetHeaders() []string {

	var tran transformation

	structype := reflect.TypeOf(tran)

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
		log.Fatalf("Failed scraping the url: %s", url)
	}

	return transformations
}
