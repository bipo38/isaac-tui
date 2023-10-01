package isaac

import (
	"strings"
)

func parseExtension(extension string) string {

	if extension == "" {
		return "Rebirth"
	}

	split := strings.Split(extension, " ")

	if split[len(split)-1] == "†" {
		return "Afterbirth Plus"
	}

	return split[len(split)-1]

}
