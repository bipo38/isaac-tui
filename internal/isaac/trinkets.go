package isaac

import (
	"isaac-scrapper/internal/utils"

	"github.com/gocolly/colly"
)

type Trinket struct {
	name, id_game, quote, effect, unlock, image string
	extension                                   utils.Extension
}

func createTrinkets(paths []Path) {

	collector := colly.NewCollector()

	var trinkets []Trinket

	collector.OnHTML(mainNode, func(h *colly.HTMLElement) {

		print(getName(h))
		// print(getId(h))
		// print(getUnlockBy(h))
		// print(getQuality(h))
		// print(getPool(h))
		// print(getExtension(h))
		// print(getItemType(h))

		// trinket := Trinket{

		// 	name: "Hola",
		// }

		// trinkets = append(trinkets, trinket)

	})
	for _, value := range paths {

		collector.Visit(string(value))

	}

	print(trinkets)

}
