package main

import (
	"fmt"
	"net"
	"bytes"
	"regexp"
	//"reflect"
)

/*
type Interface struct {
	Index        int          // positive integer that starts at one, zero is never used
	MTU          int          // maximum transmission unit
	Name         string       // e.g., "en0", "lo0", "eth0.100"
	HardwareAddr HardwareAddr // IEEE MAC-48, EUI-48 and EUI-64 form
	Flags        Flags        // e.g., FlagUp, FlagLoopback, FlagMulticast
}
*/

// getMacAddr gets the MAC hardware
// address of the host machine
func getMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	//t := reflect.TypeOf(interfaces)

	fmt.Println(interfaces)
	if err == nil {
		for _, i := range interfaces {
			fmt.Println(i)
			fmt.Println("Flags", i.Flags)
			fmt.Println("FlagUp", net.FlagUp)
			fmt.Println("HardwareAddr", i.HardwareAddr)
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				//break
			}
			fmt.Println("**********")
		}
		fmt.Println("--------------------------------------------")
	}
	return
}

func getIpAddrs() {
	fmt.Println("--------- getIpAddress")
	ipv4Regex := regexp.MustCompile(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		fmt.Println("\nName:", i.Name)
		// handle err
		for _, addr := range addrs {
			var ip string
			switch v := addr.(type) {
			case *net.IPNet:
					ip = v.IP.String()
					if ipv4Regex.MatchString(ip) {
						fmt.Println(ip)
					}
			case *net.IPAddr:
					//ip = v.IP
			}
			//fmt.Println(ip)
			// process IP address
		}
	}
	fmt.Println("*********** end of getIpAddress")
}

func main() {
	fmt.Println("Hello")
	//addr := getMacAddr()
	//fmt.Println(addr)
	getIpAddrs()
}