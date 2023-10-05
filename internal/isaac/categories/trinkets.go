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

func setTrinketImage(h *colly.HTMLElement, trinket *Trinket) error {

	imgUrl := h.ChildAttr("td:nth-child(3) a>img:nth-child(1)", "data-src")
	imgName := h.ChildAttr("td:nth-child(3) a>img:nth-child(1)", "data-image-key")

	imgPath, err := utils.DownloadImage(imgUrl, config.Trinket["imgFolder"], imgName)
	if err != nil {
		return err
	}

	trinket.image = imgPath
	return nil

}

func setTrinketUnlock(h *colly.HTMLElement, trinket *Trinket) {
	unlock := h.ChildText("div[data-source=\"unlocked by\"]>div")

	trinket.unlock = parsers.ParseUnlock(unlock)

}

func setTrinketExtension(h *colly.HTMLElement, trinket *Trinket) {
	extension := h.ChildAttr("div#context-page.context-box>img", "title")
	trinket.extension = parsers.ParseExtension(extension)
}
