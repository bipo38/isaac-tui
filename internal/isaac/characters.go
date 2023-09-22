package isaac

import (
	"fmt"
	"isaac-scrapper/config"
	"isaac-scrapper/internal/utils"

	"github.com/gocolly/colly"
)

type Character struct {
	name, unlock, image string
	extension           Extension
}

func CreateCharactersCsv() {
	var t Character

	writer, file := utils.CreateCsv(t, config.Character["csvRoute"], config.Character["csvName"])
	characters := scrapingCharacters()

	for _, v := range characters {

		character := []string{
			v.name,
			v.unlock,
			v.image,
			string(v.extension),
		}

		writer.Write(character)
	}

	defer file.Close()
	defer writer.Flush()

}

func scrapingCharacters() []Character {
	collector := colly.NewCollector()

	var characters []Character

	collector.OnHTML(config.Default["tableNode"], func(h *colly.HTMLElement) {
		character := newCharacter(h.ChildAttr("a", "href"), h)

		characters = append(characters, character)
	})

	collector.Visit(config.Character["url"])

	return characters
}

func newCharacter(path string, el *colly.HTMLElement) Character {
	character := Character{
		name:  el.ChildAttr("a", "title"),
		image: el.ChildAttr("td:nth-child(3)>a>img", "data-image-key"),
	}

	collector := colly.NewCollector()

	collector.OnHTML(config.Default["mainNode"], func(h *colly.HTMLElement) {
		setCharacterUnlock(h, &character)
		setCharacterExtension(h, &character)
		setImage(h, &character)
	})

	collector.Visit(fmt.Sprintf("%s%s", config.Default["url"], path))

	return character
}

func setImage(h *colly.HTMLElement, character *Character) {

	character.image = h.ChildAttr("img[alt=\"Character image\"]", "data-image-key")
	imgUrl := h.ChildAttr("img[alt=\"Character image\"]", "data-src")

	if imgUrl == "" {
		return
	}

	utils.DownloadImage(imgUrl, config.Character["imgRoute"], character.image)
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

	character.extension = ParseExtension(extension)
}
