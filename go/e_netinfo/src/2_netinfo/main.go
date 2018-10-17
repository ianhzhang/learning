package main

import (
	"fmt"
	"2_netinfo/netinfo"
)


func main() {
	fmt.Println("Start")

	netInfoList := netinfo.GetNetInfo()
	for _,netInfo := range netInfoList {
		fmt.Println(netInfo.Name)
		fmt.Println(netInfo)
	}
}