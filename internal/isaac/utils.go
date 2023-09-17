package isaac

import (
	"strings"
)

type Extension string

const (
	REBIRTH        Extension = "rebirth"
	AFTERBIRTH               = "afterbirth"
	AFTERBIRTHPLUS           = "afterbirthplus"
	REPENTANCE               = "repentance"
)

func ParseExtension(extension string) Extension {

	switch {
	case strings.Contains("", extension):
		return REBIRTH

	case strings.Contains("Afterbirth", extension):
		return AFTERBIRTH

	case strings.Contains("Afterbirth †", extension):
		return AFTERBIRTHPLUS

	case strings.Contains("Repentance", extension):
		return REPENTANCE

	default:
		return REBIRTH
	}

}
