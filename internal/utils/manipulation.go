package utils

import "fmt"

func routeParser(fRoute, fName string) (string, string) {
	defaultRoute := "isaac"

	route := fmt.Sprintf("%s/%s/", defaultRoute, fRoute)
	fPath := fmt.Sprintf("%s%s", route, fName)

	fmt.Println(fPath)

	return route, fPath

}
