package isaac

type Category string

const (
	ITEMS           Category = "/wiki/Items"
	TRINKETS                 = "/wiki/Trinkets"
	TRANSFORMATIONS          = "/wiki/Transformations"
	BOSSES                   = "/wiki/All_Bosses_(Bosses)"
	CHARACTERS               = "/wiki/Characters"
	CARDS                    = "/wiki/Cards_and_Runes"
	PILLS                    = "/wiki/Pills"
)

var globaLink = "https://bindingofisaacrebirth.fandom.com"
var mainNode = "div.main-container>div.resizable-container>div.has-right-rail>main.page__main>div#content>div#mw-content-text>div.mw-parser-output"
var TableNode = mainNode + ">table.wikitable>tbody>tr"
var defaultRoute = "isaac/"

// func GetPaths(category string) []string {

// 	var paths []string

// 	collector := colly.NewCollector()

// 	collector.OnHTML(TableNode, func(h *colly.HTMLElement) {
// 		path := h.ChildAttr("a", "href")

// 		if path == "" {
// 			return
// 		}

// 		print(path)

// 		paths = append(paths, globaLink+path)

// 	})

// 	collector.Visit(globaLink + category)

// 	return paths
// }
