package isaac

import (
	"isaac-scrapper/internal/system"

	"github.com/gocolly/colly"
)

type Item struct {
	name, id_game, quote, effect, unlock, image, quality, pool string
	extension                                                  Extension
}

func CreateItemsCsv() {

	var t Item

	writer, file := system.CreateCsv(t, "items", "items.csv")

	items := getItems()

	for _, v := range items {

		item := []string{
			v.name,
			v.id_game,
			v.quote,
			v.effect,
			v.unlock,
			v.image,
			v.quality,
			v.pool,
			string(v.extension),
		}

		writer.Write(item)
	}

	defer file.Close()

	defer writer.Flush()

}

func getItems() []Item {

	collector := colly.NewCollector()

	var items []Item

	collector.OnHTML(TableNode, func(h *colly.HTMLElement) {

		item := newItem(h.ChildAttr("a", "href"), h)

		if item.name == "" {
			return
		}

		items = append(items, item)
	})

	collector.Visit(globaLink + ITEMS)

	return items

}

func newItem(path string, el *colly.HTMLElement) Item {
	item := Item{
		name:      el.ChildAttr("td:nth-child(1)", "data-sort-value"),
		id_game:   el.ChildText("td:nth-child(2)"),
		quote:     el.ChildText("td:nth-child(4)"),
		effect:    el.ChildText("td:nth-child(5)"),
		image:     "imagenes3",
		quality:   el.ChildText("td:nth-child(6)"),
		extension: ParseExtension(el.ChildAttr("td:nth-child(1)>img", "title")),
	}

	collector := colly.NewCollector()

	collector.OnHTML(mainNode, func(h *colly.HTMLElement) {
		setItemUnlock(h, &item)
		setItemExtension(h, &item)
		setItemPool(h, &item)

	})

	collector.Visit(globaLink + path)

	return item

}

func setItemUnlock(h *colly.HTMLElement, item *Item) {
	unlock := h.ChildText("div[data-source=\"unlocked by\"]>div")

	if unlock != "" {
		item.unlock = unlock
	} else {
		item.unlock = "Unlocked"
	}

}

func setItemExtension(h *colly.HTMLElement, item *Item) {
	extension := h.ChildAttr("div#context-page.context-box>img", "title")

	item.extension = ParseExtension(extension)
}

func setItemPool(h *colly.HTMLElement, item *Item) {
	item.pool = h.ChildText("div[data-source=\"alias\"]>div>div.item-pool-list")
}
