package eca

import (
	"2_my_struct/common"
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
}