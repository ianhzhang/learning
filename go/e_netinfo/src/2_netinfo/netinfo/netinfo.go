package netinfo

import (
	//"fmt"
	"net"
	"bytes"
	"regexp"
)

type NetIf struct {
	Name        string
	Ipv4        string
	Mac         string
}


func GetNetInfo() []NetIf{
	if_name := ""
	if_ipv4 := ""
	if_mac := ""
	returnList := []NetIf{}
	ipv4Regex := regexp.MustCompile(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)

	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if_name = i.Name			// interface name
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				if_mac = i.HardwareAddr.String()		// interface mac addr
			}

			addrs, _ := i.Addrs()
			
			for _, addr := range addrs {
				switch v := addr.(type) {
				case *net.IPNet:
						ip := v.IP.String()
						if ipv4Regex.MatchString(ip) {
							if ip != "127.0.0.1" {
								if_ipv4 = ip	// interface ip v4 address
								netIf :=  NetIf{if_name, if_ipv4, if_mac}
								returnList = append(returnList, netIf)
							}

						}
				}
			}
		}
	}
	return returnList
}
