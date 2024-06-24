package utils

import "strings"

func RouteSplitter(route string) (firstRoute string, postRoutes string) {
	allRoutes := strings.Split(route, "/")
	firstRoute = allRoutes[0]
	postRoutes = strings.Join(allRoutes[1:], "/")
	return
}
