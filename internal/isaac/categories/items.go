package categories

import (
	"errors"
	"fmt"
	"isaac-scrapper/config"
	"isaac-scrapper/internal/isaac/parsers"
	"isaac-scrapper/internal/utils"
	"log"

	"github.com/gocolly/colly"
)

type Item struct {
	name, id_game, quote, effect, unlock, image, quality, pool, extension string
}

func CreateItemsCsv() error {

	var t Item

	writer, file, err := utils.CreateCsv(t, config.Item["csvRoute"], config.Item["csvName"])
	if err != nil {
		return err
	}

	items, err := scrapingItems()
	if err != nil {
		return err
	}

	for _, v := range items {

		err := writer.Write([]string{
			v.name,
			v.id_game,
			v.quote,
			v.effect,
			v.unlock,
			v.image,
			v.quality,
			v.pool,
			v.extension,
		})

		if err != nil {
			log.Println("error writing record to csv:", err)
			continue
		}

	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Println("csv writer error", err)
	}

	file.Close()

	return nil

}

func scrapingItems() ([]Item, error) {

	collector := colly.NewCollector()

	var items []Item

	collector.OnHTML(config.Default["tableNode"], func(el *colly.HTMLElement) {

		item, err := newItem(el)
		if err != nil {
			log.Println("error creating item:", err)
			return
		}

		items = append(items, *item)
	})

	if err := collector.Visit(config.Item["url"]); err != nil {
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

	urlPath := el.ChildAttr("a", "href")

	collector := colly.NewCollector()

	collector.OnHTML(config.Default["mainNode"], func(h *colly.HTMLElement) {

		setItemUnlock(h, &item)
		setItemExtension(h, &item)
		setItemPool(h, &item)

	})

	if err := collector.Visit(fmt.Sprintf("%s%s", config.Default["url"], urlPath)); err != nil {
		return nil, err
	}

	return &item, nil

}

func setImageItems(el *colly.HTMLElement, item *Item) error {

	imgUrl := el.ChildAttr("td:nth-child(3) a>img:nth-child(1)", "data-src")
	imgName := el.ChildAttr("td:nth-child(3) a>img:nth-child(1)", "data-image-key")

	imgPath, err := utils.DownloadImage(imgUrl, config.Item["imgFolder"], imgName)
	if err != nil {
		return err
	}

	item.image = imgPath

	return nil

}

func setItemUnlock(h *colly.HTMLElement, item *Item) {
	unlock := h.ChildText("div[data-source=\"unlocked by\"]>div")

	item.unlock = parsers.ParseUnlock(unlock)

}

func setItemExtension(h *colly.HTMLElement, item *Item) {
	extension := h.ChildAttr("div#context-page.context-box>img", "title")

	item.extension = parsers.ParseExtension(extension)
}

func setItemPool(h *colly.HTMLElement, item *Item) {
	item.pool = h.ChildText("div[data-source=\"alias\"]>div>div.item-pool-list")
}
