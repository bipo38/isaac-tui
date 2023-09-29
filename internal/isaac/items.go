package isaac

import (
	"errors"
	"fmt"
	"isaac-scrapper/config"
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

		item := []string{
			v.name,
			v.id_game,
			v.quote,
			v.effect,
			v.unlock,
			v.image,
			v.quality,
			v.pool,
			v.extension,
		}

		if err := writer.Write(item); err != nil {
			log.Println("error writing record to csv:", err)
			continue
		}
	}

	defer writer.Flush()

	defer file.Close()

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

	urlPath := el.ChildAttr("a", "href")

	item := Item{
		name:      el.ChildAttr("td:nth-child(1)", "data-sort-value"),
		id_game:   el.ChildText("td:nth-child(2)"),
		quote:     el.ChildText("td:nth-child(4)"),
		effect:    el.ChildText("td:nth-child(5)"),
		image:     "imagenes3",
		quality:   el.ChildText("td:nth-child(6)"),
		extension: parseExtension(el.ChildAttr("td:nth-child(1)>img", "title")),
	}

	if item.name == "" {
		return nil, errors.New("name is empty")
	}

	collector := colly.NewCollector()

	collector.OnHTML(config.Default["mainNode"], func(h *colly.HTMLElement) {

		setItemUnlock(h, &item)
		setItemExtension(h, &item)
		setItemPool(h, &item)

		if err := setImageItems(h, &item); err != nil {
			log.Println("error getting items image:", err)
		}
	})

	if err := collector.Visit(fmt.Sprintf("%s%s", config.Default["url"], urlPath)); err != nil {
		return nil, err
	}

	return &item, nil

}

func setItemUnlock(h *colly.HTMLElement, item *Item) {
	unlock := h.ChildText("div[data-source=\"unlocked by\"]>div")

	if unlock != "" {
		item.unlock = unlock
		return
	}

	item.unlock = "Unlocked"

}

func setItemExtension(h *colly.HTMLElement, item *Item) {
	extension := h.ChildAttr("div#context-page.context-box>img", "title")

	item.extension = parseExtension(extension)
}

func setItemPool(h *colly.HTMLElement, item *Item) {
	item.pool = h.ChildText("div[data-source=\"alias\"]>div>div.item-pool-list")
}

func setImageItems(h *colly.HTMLElement, item *Item) error {

	imgName := h.ChildAttr("img[alt=\"Item icon\"]", "data-image-key")
	imgUrl := h.ChildAttr("img[alt=\"Item icon\"]", "data-src")

	imgPath, err := utils.DownloadImage(imgUrl, config.Item["imgFolder"], imgName)
	if err != nil {
		return err
	}

	item.image = imgPath

	return nil

}
