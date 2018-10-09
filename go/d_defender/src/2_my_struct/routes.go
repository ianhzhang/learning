package main

import (
	"fmt"
	"2_my_struct/eca"
	"2_my_struct/common"
)

func GetRoutes() common.Routes {
	fmt.Println("GetRoutes")
	routes := make(common.Routes, 0)
	routes = append(routes, eca.ECARoutes...)
	return routes
}