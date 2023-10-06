package manipulation

import (
	"fmt"
	"isaac-scrapper/config"
	"reflect"
	"strings"
)

func RouteParser(fRoute, fName string) (string, string) {
	s := config.Default["routeStart"]

	r := fmt.Sprintf("%s/%s/", s, fRoute)
	p := fmt.Sprintf("%s%s", r, fName)

	return r, p

}

func ParserFileName(fName, ext string) string {

	if strings.Contains(ext, ".") {
		return fmt.Sprintf("%s%s", fName, ext)
	}

	return fmt.Sprintf("%s.%s", fName, ext)

}

func GetHeaders(el interface{}) []string {

	t := reflect.TypeOf(el)

	var headers []string

	for i := 0; i < t.NumField(); i++ {

		headers = append(headers, t.Field(i).Name)
	}

	return headers
}
