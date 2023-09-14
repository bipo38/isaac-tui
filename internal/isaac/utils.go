package isaac

import "reflect"

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

func GetHeaders[C any](t C) []string {
	structype := reflect.TypeOf(t)

	var headers []string

	for i := 0; i < structype.NumField(); i++ {

		headers = append(headers, structype.Field(i).Name)
	}

	return headers
}
