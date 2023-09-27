package utils

import (
	"fmt"
	"isaac-scrapper/config"
	"reflect"
	"strings"
)

func RouteParser(fRoute, fName string) (string, string) {
	startRoute := config.Default["folderStart"]

	route := fmt.Sprintf("%s/%s/", startRoute, fRoute)
	fPath := fmt.Sprintf("%s%s", route, fName)

	return route, fPath

}

func ParserFileName(fName, extension string) string {

	if strings.Contains(extension, ".") {
		return fmt.Sprintf("%s%s", fName, extension)
	}

	return fmt.Sprintf("%s.%s", fName, extension)

}

func GetHeaders(element interface{}) []string {

	structype := reflect.TypeOf(element)

	var headers []string

	for i := 0; i < structype.NumField(); i++ {

		headers = append(headers, structype.Field(i).Name)
	}

	return headers
}
