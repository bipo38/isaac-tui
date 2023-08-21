package trinket

import "isaac-scrapper/internal/utils"

type trinket struct {
	name, id_game, pickup, effect, unlock string
	extension                             utils.Extension
}

func NewTrinket(name, id_game, pickup, effect, unlock string, extension utils.Extension) trinket {

	trinket := trinket{
		name:      name,
		id_game:   id_game,
		pickup:    pickup,
		effect:    effect,
		unlock:    unlock,
		extension: extension,
	}

	return trinket
}
