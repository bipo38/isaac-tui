package isaac

import (
	"encoding/csv"
	"isaac-scrapper/internal/system"
	"strings"

	"github.com/gocolly/colly"
)

type Pickup struct {
	name, id_game, quote, effect, unlock, image, kind string
	extension                                         Extension
}

func CreatePickupsCsv() {

	var t Pickup

	fName := "pickups.csv"
	route := defaultRoute + "pickups/"
	fullRoute := route + fName

	pickups := scrapingPickups()
	headers := GetHeaders(t)
	system.CreateDirs(route)
	file := system.CreateFile(fullRoute)

	writer := csv.NewWriter(file)
	writer.Write(headers)

	for _, v := range pickups {

		pickup := []string{
			v.name,
			v.id_game,
			v.quote,
			v.effect,
			v.unlock,
			v.image,
			v.kind,
			string(v.extension),
		}

		writer.Write(pickup)
	}

	defer file.Close()

	defer writer.Flush()

}

func scrapingPickups() []Pickup {

	collector := colly.NewCollector()

	var pickups []Pickup

	collector.OnHTML(TableNode, func(h *colly.HTMLElement) {
		pickup := newPickup(h, len(pickups))

		if pickup.name == "" || strings.Contains(pickup.name, "https") {
			return
		}

		pickups = append(pickups, pickup)
	})

	collector.Visit(globaLink + CARDS)

	return pickups
}

func newPickup(el *colly.HTMLElement, len int) Pickup {

	quote := "4"
	effect := "5"

	kind := "card"
	name := el.ChildAttr("td:nth-child(1)", "data-sort-value")

	if name == "Rune of Hagalaz " {
		kind = "rune"
	}

	if name == "Dice Shard" {
		kind = "other"
	}

	pickup := Pickup{
		name:      name,
		id_game:   el.ChildText("td:nth-child(2)"),
		effect:    el.ChildText("td:nth-child(" + effect + ")"),
		image:     "image",
		quote:     el.ChildText("td:nth-child(" + quote + ")"),
		unlock:    "unlocked",
		extension: ParseExtension(el.ChildAttr("td:nth-child(1)>img", "title")),
		kind:      kind,
	}

	if len > 22 {
		pickup.unlock = el.ChildText("td:nth-child(4)")
		quote = "5"
		effect = "6"

	}

	return pickup
}
