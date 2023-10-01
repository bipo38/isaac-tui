package isaac

import (
	"errors"
	"fmt"
	"isaac-scrapper/config"
	"isaac-scrapper/internal/utils"
	"log"

	"github.com/gocolly/colly"
)

type Character struct {
	name, unlock, image, extension string
}

func CreateCharactersCsv() error {
	var t Character

	writer, file, err := utils.CreateCsv(t, config.Character["csvRoute"], config.Character["csvName"])
	if err != nil {
		return err
	}

	characters, err := scrapingCharacters()
	if err != nil {
		return err
	}

	for _, v := range characters {

		character := []string{
			v.name,
			v.unlock,
			v.image,
			v.extension,
		}

		if err := writer.Write(character); err != nil {
			log.Printf("error writing character: %v", err)
			continue
		}

	}

	defer file.Close()

	defer writer.Flush()

	return nil
}

func scrapingCharacters() ([]Character, error) {
	collector := colly.NewCollector()

	var characters []Character

	collector.OnHTML(config.Default["tableNode"], func(el *colly.HTMLElement) {

		character, err := newCharacter(el)
		if err != nil {
			log.Printf("error creating character: %v", err)
			return
		}

		characters = append(characters, *character)
	})

	if err := collector.Visit(config.Character["url"]); err != nil {
		return nil, err
	}

	return characters, nil
}

func newCharacter(el *colly.HTMLElement) (*Character, error) {

	urlPath := el.ChildAttr("a", "href")

	character := Character{
		name:  el.ChildAttr("a", "title"),
		image: el.ChildAttr("td:nth-child(3)>a>img", "data-image-key"),
	}

	if character.name == "" {
		return nil, errors.New("name is empty")
	}

	collector := colly.NewCollector()

	collector.OnHTML(config.Default["mainNode"], func(h *colly.HTMLElement) {
		setCharacterUnlock(h, &character)
		setCharacterExtension(h, &character)

		if err := setImageCharacters(h, &character); err != nil {
			log.Printf("error getting character image: %v", err)
		}
	})

	if err := collector.Visit(fmt.Sprintf("%s%s", config.Default["url"], urlPath)); err != nil {
		return nil, err
	}

	return &character, nil
}

func setCharacterExtension(h *colly.HTMLElement, character *Character) {
	extension := h.ChildAttr("div#context-page.context-box>img", "title")

	character.extension = parseExtension(extension)
}

func setCharacterUnlock(h *colly.HTMLElement, character *Character) {
	unlock := h.ChildText("div[data-source=\"unlocked by\"]>div")

	character.unlock = isUnlock(unlock)

}

func setImageCharacters(h *colly.HTMLElement, character *Character) error {

	imgName := h.ChildAttr("img[alt=\"Character image\"]", "data-image-key")
	imgUrl := h.ChildAttr("img[alt=\"Character image\"]", "data-src")

	imgPath, err := utils.DownloadImage(imgUrl, config.Character["imgFolder"], imgName)
	if err != nil {
		return err
	}

	character.image = imgPath

	return nil

}
