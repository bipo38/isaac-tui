package isaac

import (
	"isaac-scrapper/internal/utils"

	"github.com/gocolly/colly"
)

func getName(h *colly.HTMLElement) string {
	return h.ChildText("table.wikitable>tbody>tr>td>a")
}

func getId(h *colly.HTMLElement) string {
	return h.ChildText("div[data-source=\"id\"]>div>code")
}

func getUnlockBy(h *colly.HTMLElement) string {
	unlock := h.ChildText("div[data-source=\"unlocked by\"]>div")

	if unlock == "" {
		return "Unlocked"
	}
	return unlock

}

func getQuote(h *colly.HTMLElement) string {

	return h.ChildText("div[data-source=\"quote\"]>div")
}

func getQuality(h *colly.HTMLElement) string {
	return h.ChildText("div[data-source=\"quality\"]>div")
}

func getPool(h *colly.HTMLElement) []string {

	pools := h.ChildAttrs("div[data-source=\"alias\"]>div>div.item-pool-list>span[style=\"font-size: smaller; font-weight: bold\"]>a", "title")

	if len(pools) == 0 {
		return append(pools, "None")
	}

	return pools

}

// func getEffect(h *colly.HTMLElement) string {

// 	return h.ChildText()
// }

// func getImage(h *colly.HTMLElement) string {

// }

func getExtension(h *colly.HTMLElement) string {
	extension := h.ChildAttr("div#context-page.context-box>img", "title")

	return string(utils.ParseExtension(extension))
}

func getItemType(h *colly.HTMLElement) string {
	return h.ChildText("p>a[title=\"Items\"]")
}
