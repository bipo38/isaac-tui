package categories

import (
	"errors"
	"fmt"

	"github.com/bipo38/isaac-tui/internal/downloads"

	"log"

	"github.com/bipo38/isaac-tui/config"
	"github.com/bipo38/isaac-tui/internal/creates"
	"github.com/bipo38/isaac-tui/internal/isaac/parsers"
	"github.com/gocolly/colly"
)

type Item struct {
	name, id_game, quote, effect, unlock, image, quality, pool, extension string
}

func CreateItemsCsv() error {

	items, err := scrapingItems()
	if err != nil {
		return err
	}

	csv := creates.Csv[Item]{
		Name:     config.Item["csvName"],
		Path:     config.Item["csvRoute"],
		Category: items,
	}

	if err := csv.Write(); err != nil {
		return err
	}

	return nil

}

func scrapingItems() ([]Item, error) {

	c := colly.NewCollector()

	var items []Item

	c.OnHTML(config.Default["tableNode"], func(el *colly.HTMLElement) {

		i, err := newItem(el)
		if err != nil {
			log.Println("error creating item:", err)
			return
		}

		items = append(items, *i)
	})

	if err := c.Visit(config.Item["url"]); err != nil {
		return nil, err
	}

	return items, nil

}

func newItem(el *colly.HTMLElement) (*Item, error) {

	item := Item{
		name: el.ChildAttr("td:nth-child(1)", "data-sort-value"),
	}

	if item.name == "" || item.name == "Tonsil" {
		return nil, errors.New("name is empty")
	}

	item.id_game = el.ChildText("td:nth-child(2)")
	item.quote = el.ChildText("td:nth-child(4)")
	item.effect = el.ChildText("td:nth-child(5)")
	item.quality = el.ChildText("td:nth-child(6)")

	if err := setImageItems(el, &item); err != nil {
		return nil, err
	}

	path := el.ChildAttr("a", "href")

	c := colly.NewCollector()

	c.OnHTML(config.Default["mainNode"], func(h *colly.HTMLElement) {

		setItemUnlock(h, &item)
		setItemExtension(h, &item)
		setItemPool(h, &item)

	})

	if err := c.Visit(fmt.Sprintf("%s%s", config.Default["url"], path)); err != nil {
		return nil, err
	}

	return &item, nil

}

func setImageItems(el *colly.HTMLElement, item *Item) error {

	url := el.ChildAttr("td:nth-child(3) a>img:nth-child(1)", "data-src")
	name := el.ChildAttr("td:nth-child(3) a>img:nth-child(1)", "data-image-key")

	p, err := downloads.Image(url, config.Item["imgFolder"], name)
	if err != nil {
		return err
	}

	item.image = p

	return nil

}

func setItemUnlock(h *colly.HTMLElement, item *Item) {
	u := h.ChildText("div[data-source=\"unlocked by\"]>div")

	item.unlock = parsers.Unlock(u)

}

func setItemExtension(h *colly.HTMLElement, item *Item) {
	e := h.ChildAttr("div#context-page.context-box>img", "title")

	item.extension = parsers.Extension(e)
}

func setItemPool(h *colly.HTMLElement, item *Item) {
	item.pool = h.ChildText("div[data-source=\"alias\"]>div>div.item-pool-list")
}
