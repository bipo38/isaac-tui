package isaac

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Category string

var print = fmt.Println

const (
	ITEMS           Category = "Items"
	TRINEKTS                 = "Trinkets"
	TRANSFORMATIONS          = "Transformations"
	BOSSES                   = "All_Bosses_(Bosses)"
)

type Element string

const (
	TABLE     Element = ">table.wikitable>tbody>tr"
	BOSS_PAGE         = ">div.table-wide>div.table-wide-inner>table>tbody>tr>td>div>div"
	PAGE              = ""
)

func DoScraping(page Category, el Element) {

	linkPage := "https://bindingofisaacrebirth.fandom.com/wiki/" + string(page)

	node := "div.main-container>div.resizable-container>div.has-right-rail>main.page__main>div#content>div#mw-content-text>div.mw-parser-output" + string(el)

	print(linkPage)
	print(node)
	// var elements []utils.Base

	collector := colly.NewCollector()

	collector.OnHTML(node, func(h *colly.HTMLElement) {

		url := h.ChildAttr("a", "href")

		print(url)

		// coll := colly.NewCollector()
		// hola(&coll)
		// element := utils.Base{
		// 	Name:    h.ChildAttr("td:nth-child(1)", "data-sort-value"),
		// 	Id_game: h.ChildText("td:nth-child(2)"),
		// 	// quote:     h.ChildText("td:nth-child(4)"),
		// 	Effect: h.ChildText("td:nth-child(5)"),
		// 	// image:     "imagenes3",
		// 	// quality:   h.ChildText("td:nth-child(6)"),
		// 	Extension: utils.ParseExtension(h.ChildAttr("td:nth-child(1)>img", "title")),
		// }

		// elements = append(elements, element)

	})

	collector.Visit(linkPage)

	// return elements
}

// func hola(**colly.Collector) {

// }
