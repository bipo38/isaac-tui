package utils

import "isaac-scrapper/internal/utils"

type Extension string

const (
	REBIRTH        Extension = "rebirth"
	AFTERBIRTH               = "afterbirth"
	AFTERBIRTHPLUS           = "afterbirthplus"
	REPENTANCE               = "repentance"
)

func ParseExtension(extension string) Extension {

	switch extension {
	case "Added in Afterbirth":
		return AFTERBIRTH

	case "Added in Afterbirth †":
		return AFTERBIRTHPLUS

	case "Added in Repentance":
		return REPENTANCE

	default:
		return REBIRTH
	}

}

type Base struct {
	Name      string          `base:"name"`
	Id_game   string          `base:"id_game"`
	Effect    string          `base:"effect"`
	Unlock    string          `base:"unlock"`
	Extension utils.Extension `base:"extension"`
	Image     string          `base:"image"`
}

//Item ->
//Transformation ->
//Trinket -> quote,
