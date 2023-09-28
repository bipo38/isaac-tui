package isaac

import (
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
			string(v.extension),
		}

		if err := writer.Write(character); err != nil {
			log.Printf("error writing character: %v", err)
			continue
		}

	}

	defer writer.Flush()

	defer file.Close()

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

	collector := colly.NewCollector()

	collector.OnHTML(config.Default["mainNode"], func(h *colly.HTMLElement) {
		setCharacterUnlock(h, &character)
		setCharacterExtension(h, &character)

		if err := setImage(h, &character); err != nil {
			log.Printf("error setting image: %v", err)
			character.image = "Error Downloading Image"

		}
	})

	if err := collector.Visit(fmt.Sprintf("%s%s", config.Default["url"], urlPath)); err != nil {
		return nil, err
	}

	return &character, nil
}

func setImage(h *colly.HTMLElement, character *Character) error {

	character.image = h.ChildAttr("img[alt=\"Character image\"]", "data-image-key")
	imgUrl := h.ChildAttr("img[alt=\"Character image\"]", "data-src")

	if imgUrl == "" {
		return nil
	}

	if err := utils.DownloadImage(imgUrl, config.Character["imgRoute"], character.image); err != nil {
		return err
	}

	return nil

}

func setCharacterUnlock(h *colly.HTMLElement, character *Character) {
	unlock := h.ChildText("div[data-source=\"unlocked by\"]>div")

	if unlock != "" {
		character.unlock = unlock
	} else {
		character.unlock = "Unlocked"
	}

}

func setCharacterExtension(h *colly.HTMLElement, character *Character) {
	extension := h.ChildAttr("div#context-page.context-box>img", "title")

	character.extension = parseExtension(extension)
}
