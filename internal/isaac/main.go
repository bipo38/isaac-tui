package isaac

import (
	"fmt"
)

type Category string
type Path string

var print = fmt.Println

var globaLink = "https://bindingofisaacrebirth.fandom.com/wiki"
var mainNode = "div.main-container>div.resizable-container>div.has-right-rail>main.page__main>div#content>div#mw-content-text>div.mw-parser-output"
var TableNode = mainNode + ">table.wikitable>tbody>tr"

const (
	ITEMS           Category = "/Items"
	TRINKETS                 = "/Trinkets"
	TRANSFORMATIONS          = "/Transformations"
	BOSSES                   = "/All_Bosses_(Bosses)"
	CHARACTERS               = "/Characters"
	CARDS                    = "/Cards_and_Runes"
	PILLS                    = "/Pills"
)

type Section string

// const (
// TABLE Section = ">table.wikitable>tbody>tr"
// BOSS_PAGE         = ">div.table-wide>div.table-wide-inner>table"
// PAGE = ""
// )

func StartScraping(category Category) {

	url := globaLink + string(category)
	GetPaths(url)
	// GetQuote(url)
	// GetID(url, category)
	// createTrinkets(GetPaths(url))

	// print(getBaseElement(url))

}
