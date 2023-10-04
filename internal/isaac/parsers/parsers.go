package parsers

import (
	"strings"
)

func ParseExtension(extension string) string {

	if extension == "" {
		return "Rebirth"
	}

	split := strings.Split(extension, " ")

	if split[len(split)-1] == "†" {
		return "Afterbirth Plus"
	}

	return split[len(split)-1]

}

func ParseUnlock(unlock string) string {

	if unlock == "" {
		return "Unlocked"
	}

	return unlock
}
