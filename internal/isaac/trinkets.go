package isaac

import (
	"fmt"
	"isaac-scrapper/config"
	"isaac-scrapper/internal/utils"
	"log"

	"github.com/gocolly/colly"
)

type Trinket struct {
	name, id_game, quote, effect, unlock, image, extension string
}

func CreateTrinketsCsv() error {

	var t Trinket

	writer, file, err := utils.CreateCsv(t, config.Trinket["csvRoute"], config.Trinket["csvName"])
	if err != nil {
		return err
	}

	trinkets, err := scrapingTrinkets()
	if err != nil {
		return err
	}

	for _, v := range trinkets {

		trinket := []string{
			v.name,
			v.id_game,
			v.quote,
			v.effect,
			v.unlock,
			v.image,
			v.extension,
		}

		if err := writer.Write(trinket); err != nil {
			continue
		}

	}

	defer file.Close()

	defer writer.Flush()

	return nil

}

func scrapingTrinkets() ([]Trinket, error) {

	collector := colly.NewCollector()

	var trinkets []Trinket

	collector.OnHTML(config.Default["tableNode"], func(el *colly.HTMLElement) {

		trinket, err := newTrinket(el)
		if err != nil {
			log.Printf("error creating trinket: %v", err)
			return
		}

		trinkets = append(trinkets, *trinket)
	})

	if err := collector.Visit(config.Trinket["url"]); err != nil {
		return nil, err
	}

	return trinkets, nil

}

func newTrinket(el *colly.HTMLElement) (*Trinket, error) {
	urlPath := el.ChildAttr("a", "href")

	trinket := Trinket{
		name:    el.ChildAttr("td:nth-child(1)", "data-sort-value"),
		id_game: el.ChildText("td:nth-child(2)"),
		quote:   el.ChildText("td:nth-child(4)"),
		effect:  el.ChildText("td:nth-child(5)"),
		image:   "image",
	}

	collector := colly.NewCollector()

	collector.OnHTML(config.Default["mainNode"], func(h *colly.HTMLElement) {
		setTrinketUnlock(h, &trinket)
		setTrinketExtension(h, &trinket)

	})

	if err := collector.Visit(fmt.Sprintf("%s%s", config.Default["url"], urlPath)); err != nil {
		return nil, err
	}

	return &trinket, nil

}

func setTrinketUnlock(h *colly.HTMLElement, trinket *Trinket) {
	unlock := h.ChildText("div[data-source=\"unlocked by\"]>div")

	if unlock != "" {
		trinket.unlock = unlock
	} else {
		trinket.unlock = "Unlocked"
	}

}

func setTrinketExtension(h *colly.HTMLElement, trinket *Trinket) {
	extension := h.ChildAttr("div#context-page.context-box>img", "title")
	trinket.extension = parseExtension(extension)
}
