package utils

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
