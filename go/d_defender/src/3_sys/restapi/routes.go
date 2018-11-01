package restapi

import (
	"fmt"

	"3_sys/restapi/common"
	"3_sys/restapi/eca"
)

func GetRoutes() common.Routes {
	fmt.Println("GetRoutes")
	routes := make(common.Routes, 0)
	routes = append(routes, eca.ECARoutes...)
	return routes
}