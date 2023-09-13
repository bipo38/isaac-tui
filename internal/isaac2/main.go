package isaac2

import (
	"sync"

	"github.com/gocolly/colly"
)

type Base interface {
	setName(h *colly.HTMLElement)
}

type Category string

const (
	ITEMS           Category = "/Items"
	TRINKETS                 = "/Trinkets"
	TRANSFORMATIONS          = "/Transformations"
	BOSSES                   = "/All_Bosses_(Bosses)"
	CHARACTERS               = "/Characters"
	CARDS                    = "/Cards_and_Runes"
	PILLS                    = "/Pills"
)

var globaLink = "https://bindingofisaacrebirth.fandom.com/wiki"
var mainNode = "div.main-container>div.resizable-container>div.has-right-rail>main.page__main>div#content>div#mw-content-text>div.mw-parser-output"
var TableNode = mainNode + ">table.wikitable>tbody>tr"

var lock = &sync.Mutex{}

type scrapper struct {
	collector *colly.Collector
}

var scrapperInstance *scrapper

func getScrapperInstance() *scrapper {
	if scrapperInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if scrapperInstance == nil {
			scrapperInstance = &scrapper{}
		}
	}

	return scrapperInstance
}
