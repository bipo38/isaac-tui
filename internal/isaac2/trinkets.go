package isaac2

import (
	"fmt"
	"isaac-scrapper/internal/utils"

	"github.com/gocolly/colly"
)

type Table struct {
	name, quote, effect, id_game string
}

type Trinket struct {
	unlock    string
	extension utils.Extension
	Table
}

func GetTrinkets() []Trinket {

	collector := colly.NewCollector()

	var trinkets []Trinket

	collector.OnHTML(mainNode, func(h *colly.HTMLElement) {

		table := Table{
			name:    h.ChildText("table.wikitable>tbody>tr>td>a"),
			id_game: h.ChildText("td:nth-child(2)"),
			quote:   h.ChildText("td:nth-child(4)"),
			effect:  h.ChildText("td:nth-child(5)"),
		}

		trinkets = append(trinkets, newTrinket(table, h.ChildAttr("a", "href")))
		fmt.Println(table)
	})

	collector.Visit(globaLink + string(TRINKETS))

	return trinkets

}

func newTrinket(table Table, path string) Trinket {

	collector := colly.NewCollector()

	trinket := Trinket{
		Table: table,
	}

	collector.OnHTML(mainNode, func(h *colly.HTMLElement) {
		
	})

	collector.Visit(globaLink + path)

	return trinket
}

func (trinket *Trinket) setName(h *colly.HTMLElement) string {
	return "H"
}

// func getPath(h *colly.HTMLElement) string {

// }
