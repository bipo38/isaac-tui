package categories

import (
	"errors"
	"fmt"

	"log"

	"github.com/bipo38/isaac-tui/config"
	"github.com/bipo38/isaac-tui/internal/creates"
	"github.com/bipo38/isaac-tui/internal/downloads"
	"github.com/bipo38/isaac-tui/internal/isaac/parsers"
	"github.com/gocolly/colly"
)

type Trinket struct {
	name, id_game, quote, effect, unlock, image, extension string
}

func CreateTrinketsCsv() error {

	trinkets, err := scrapingTrinkets()
	if err != nil {
		return err
	}

	csv := creates.Csv[Trinket]{
		Name:     config.Trinket["csvName"],
		Path:     config.Trinket["csvRoute"],
		Category: trinkets,
	}

	if err := csv.Write(); err != nil {
		return err
	}

	return nil

}

func scrapingTrinkets() ([]Trinket, error) {

	c := colly.NewCollector()

	var trinkets []Trinket

	c.OnHTML(config.Default["tableNode"], func(el *colly.HTMLElement) {

		t, err := newTrinket(el)
		if err != nil {
			log.Printf("error creating trinket: %v", err)
			return
		}

		trinkets = append(trinkets, *t)
	})

	if err := c.Visit(config.Trinket["url"]); err != nil {
		return nil, err
	}

	return trinkets, nil

}

func newTrinket(el *colly.HTMLElement) (*Trinket, error) {

	trinket := Trinket{
		name: el.ChildAttr("td:nth-child(1)", "data-sort-value"),
	}

	if trinket.name == "" {
		return nil, errors.New("name is empty")
	}

	trinket.id_game = el.ChildText("td:nth-child(2)")
	trinket.quote = el.ChildText("td:nth-child(4)")
	trinket.effect = el.ChildText("td:nth-child(5)")

	if err := setTrinketImage(el, &trinket); err != nil {
		log.Printf("error getting trinket image: %v", err)
	}

	path := el.ChildAttr("a", "href")

	c := colly.NewCollector()

	c.OnHTML(config.Default["mainNode"], func(h *colly.HTMLElement) {
		setTrinketUnlock(h, &trinket)
		setTrinketExtension(h, &trinket)

	})

	if err := c.Visit(fmt.Sprintf("%s%s", config.Default["url"], path)); err != nil {
		return nil, err
	}

	return &trinket, nil

}

func setTrinketImage(h *colly.HTMLElement, trinket *Trinket) error {

	url := h.ChildAttr("td:nth-child(3) a>img:nth-child(1)", "data-src")
	name := h.ChildAttr("td:nth-child(3) a>img:nth-child(1)", "data-image-key")

	p, err := downloads.Image(url, config.Trinket["imgFolder"], name)
	if err != nil {
		return err
	}

	trinket.image = p

	return nil
}

func setTrinketUnlock(h *colly.HTMLElement, trinket *Trinket) {
	u := h.ChildText("div[data-source=\"unlocked by\"]>div")

	trinket.unlock = parsers.Unlock(u)

}

func setTrinketExtension(h *colly.HTMLElement, trinket *Trinket) {
	e := h.ChildAttr("div#context-page.context-box>img", "title")
	trinket.extension = parsers.Extension(e)
}
