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

type Character struct {
	name, unlock, image, extension string
}

func CreateCharactersCsv() error {

	characters, err := scrapingCharacters()
	if err != nil {
		return err
	}

	csv := creates.Csv[Character]{
		Name:     config.Character["csvName"],
		Path:     config.Character["csvRoute"],
		Category: characters,
	}

	if err := csv.Write(); err != nil {
		return err
	}

	return nil

}

func scrapingCharacters() ([]Character, error) {
	c := colly.NewCollector()

	var characters []Character

	c.OnHTML(config.Default["tableNode"], func(el *colly.HTMLElement) {

		ch, err := newCharacter(el)
		if err != nil {
			log.Printf("error creating character: %v", err)
			return
		}

		characters = append(characters, *ch)
	})

	if err := c.Visit(config.Character["url"]); err != nil {
		return nil, err
	}

	return characters, nil
}

func newCharacter(el *colly.HTMLElement) (*Character, error) {

	character := Character{
		name: el.ChildAttr("a", "title"),
	}

	if character.name == "" {
		return nil, errors.New("name is empty")
	}

	path := el.ChildAttr("a", "href")

	c := colly.NewCollector()

	c.OnHTML(config.Default["mainNode"], func(h *colly.HTMLElement) {
		setCharacterUnlock(h, &character)
		setCharacterExtension(h, &character)

		if err := setImageCharacters(h, &character); err != nil {
			log.Printf("error getting character image: %v", err)
		}
	})

	if err := c.Visit(fmt.Sprintf("%s%s", config.Default["url"], path)); err != nil {
		return nil, err
	}

	return &character, nil
}

func setCharacterUnlock(h *colly.HTMLElement, character *Character) {
	u := h.ChildText("div[data-source=\"unlocked by\"]>div")

	if u == "" {
		u = h.ChildText("div.infobox2>div:last-child")
	}

	character.unlock = parsers.Unlock(u)

}

func setCharacterExtension(h *colly.HTMLElement, character *Character) {
	ext := h.ChildAttr("div#context-page.context-box>img", "title")

	character.extension = parsers.Extension(ext)
}

func setImageCharacters(h *colly.HTMLElement, character *Character) error {

	name := h.ChildAttr("img[alt=\"Character image\"]", "data-image-key")
	url := h.ChildAttr("img[alt=\"Character image\"]", "data-src")

	if name == "" {

		name = h.ChildAttr("div.infobox2>div:nth-child(2) img:nth-child(1)", "data-image-key")
		url = h.ChildAttr("div.infobox2>div:nth-child(2) img:nth-child(1)", "data-src")
	}

	p, err := downloads.Image(url, config.Character["imgFolder"], name)
	if err != nil {
		return err
	}

	character.image = p

	return nil

}
