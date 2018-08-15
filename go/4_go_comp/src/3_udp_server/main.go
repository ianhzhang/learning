package main

import (
	"fmt"
	"net"
)


func main() {
	fmt.Println("UDP Server!")
	addr := net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", &addr) // code does not block here
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	
	var buf [1024]byte
	for {
		rLen, remoteAddr, err := conn.ReadFromUDP(buf[:])
		// Do stuff with the read bytes
		if err!= nil {
			fmt.Println(err)
			panic(err)
		} else {
			fmt.Println("recv length:", rLen)
			fmt.Println("remote:", remoteAddr)
			fmt.Println("recv buf:", buf)
			fmt.Println("recv buf:", string(buf[:1024]))
		}
	}
}