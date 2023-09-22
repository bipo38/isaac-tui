package utils

import (
	"fmt"
	"isaac-scrapper/config"
)

func routeParser(fRoute, fName string) (string, string) {
	defaultRoute := config.Default["folderDefaultRoute"]

	route := fmt.Sprintf("%s/%s/", defaultRoute, fRoute)
	fPath := fmt.Sprintf("%s%s", route, fName)

	return route, fPath

}
