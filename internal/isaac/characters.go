package isaac

import (
	"fmt"
	"isaac-scrapper/internal/utils"
	"log"

	"github.com/gocolly/colly"
)

type Character struct {
	name, unlock, image string
	extension           Extension
}

func CreateCharactersCsv() {
	var t Character

	writer, file := utils.CreateCsv(t, "characters", "characters.csv")
	characters := scrapingCharacters()

	for _, v := range characters {

		character := []string{
			v.name,
			v.unlock,
			v.image,
			string(v.extension),
		}

		if err := writer.Write(character); err != nil {
			fmt.Println(fmt.Errorf("Falied writing character in csv field: %s", v.name))
			continue
		}

	}

	defer file.Close()
	defer writer.Flush()

}

func scrapingCharacters() []Character {
	collector := colly.NewCollector()

	var characters []Character

	collector.OnHTML(TableNode, func(h *colly.HTMLElement) {
		character := newCharacter(h.ChildAttr("a", "href"), h)

		characters = append(characters, character)
	})

	if err := collector.Visit(globaLink + CHARACTERS); err != nil {

		log.Fatalf("Failed to start scraping: %s ", err , )
	}

	return characters
}

func newCharacter(path string, el *colly.HTMLElement) Character {
	character := Character{
		name:  el.ChildAttr("a", "title"),
		image: el.ChildAttr("td:nth-child(3)>a>img", "data-image-key"),
	}

	collector := colly.NewCollector()

	collector.OnHTML(mainNode, func(h *colly.HTMLElement) {
		setCharacterUnlock(h, &character)
		setCharacterExtension(h, &character)
		setImage(h, &character)
	})

	collector.Visit(globaLink + path)

	return character
}

func setImage(h *colly.HTMLElement, character *Character) {

	character.image = h.ChildAttr("img[alt=\"Character image\"]", "data-image-key")
	imgUrl := h.ChildAttr("img[alt=\"Character image\"]", "data-src")

	if imgUrl == "" {
		return
	}

	utils.DownloadImage(imgUrl, "characters/images", character.image)
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
