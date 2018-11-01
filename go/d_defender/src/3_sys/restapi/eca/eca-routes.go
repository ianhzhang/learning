package eca

import (
	"3_sys/restapi/common"
)

var ECARoutes = []common.Route{
	{
		"GetHello1",
		"GET",
		"/eca/getHello1",
		GetHello1,
	},
	{
		"GetHello22",
		"GET",
		"/eca/getHello2",
		GetHello2,
	},
	{
		"GetData",
		"GET",
		"/eca/getData/{name}/{seq}",
		GetData,
	},
}