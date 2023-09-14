package isaac

import "github.com/gocolly/colly"

type Base struct {
	name, description, path string
}

func getPaths(url string) []Path {

	var paths []Path

	collector := colly.NewCollector()

	collector.OnHTML(TableNode, func(h *colly.HTMLElement) {
		path := h.ChildAttr("a", "href")

		if path == "" {
			return
		}

		print(path)

		paths = append(paths, Path(globaLink+path))

	})

	collector.Visit(url)

	return paths
}

func GetQuote(url string) []Path {

	var paths []Path

	collector := colly.NewCollector()

	collector.OnHTML(TableNode, func(h *colly.HTMLElement) {
		path := h.ChildText("td>i")

		if path == "" {
			return
		}

		print(path)

		paths = append(paths, Path(globaLink+path))

	})

	collector.Visit(url)

	return paths
}

func getBaseElement(url string) []Base {

	var elements []Base

	collector := colly.NewCollector()

	collector.OnHTML(TableNode, func(h *colly.HTMLElement) {

		element := Base{
			name: h.ChildAttr("a", "title"),
			// description: ,
			path: h.ChildAttr("a", "href"),
		}

		elements = append(elements, element)

	})

	collector.Visit(url)

	return elements

}

func GetDescription(url string, category Category) []Path {

	var paths []Path
	var child string

	if category == TRANSFORMATIONS {
		child = "4"
	} else {
		child = "5"
	}

	collector := colly.NewCollector()

	collector.OnHTML(TableNode, func(h *colly.HTMLElement) {
		if len(paths) > 21 {
			child = "6"
		}

		path := h.ChildText("td:nth-child(" + child + ")")

		if path == "" {
			return
		}
		print(path)

		paths = append(paths, Path(globaLink+path))

	})

	collector.Visit(url)
	return paths
}

func GetID(url string, category Category) []Path {

	var paths []Path

	var child string

	if category == TRANSFORMATIONS || category == PILLS {
		child = "1"
	} else {
		child = "2"
	}

	collector := colly.NewCollector()

	collector.OnHTML(TableNode, func(h *colly.HTMLElement) {
		path := h.ChildText("td:nth-child(" + child + ")")
		// if len(paths) > 719 {
		// 	child = "1"
		// }

		if path == "" {
			return
		}

		print(path)

		paths = append(paths, Path(globaLink+path))

	})

	collector.Visit(url)

	return paths
}
