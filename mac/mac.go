package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("get mac address")
	ifas, _ := net.Interfaces()
	for _, ifa := range ifas {
		if strings.HasPrefix(ifa.Name, "en") {
			fmt.Println(ifa.HardwareAddr)
		}
	}
}