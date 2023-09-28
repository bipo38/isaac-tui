package isaac

import (
	"strings"
)

func parseExtension(extension string) string {

	if extension == "" {
		return "Rebirth"
	}

	if extension == "Added in Afterbirth †" {
		return "Afterbirth Plus"
	}

	return strings.ReplaceAll(extension, "Added in", "")

}
