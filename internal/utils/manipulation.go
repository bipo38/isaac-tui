package utils

import (
	"fmt"
	"isaac-scrapper/config"
	"reflect"
	"strings"
)

func RouteParser(fRoute, fName string) (string, string) {
	routeStart := config.Default["routeStart"]

	route := fmt.Sprintf("%s/%s/", routeStart, fRoute)
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

	structType := reflect.TypeOf(element)

	var headers []string

	for i := 0; i < structType.NumField(); i++ {

		headers = append(headers, structType.Field(i).Name)
	}

	return headers
}
