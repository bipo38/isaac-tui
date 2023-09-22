package isaac

import (
	"strings"
)

func ParseExtension(extension string) string {

	if "" == extension {
		return "Rebirth"
	}

	if "Added in Afterbirth †" == extension {
		return "Afterbirth Plus"
	}

	return strings.ReplaceAll(extension, "Added in", "")

}
