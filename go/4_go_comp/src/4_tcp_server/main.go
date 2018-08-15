package main

import (
	"fmt"
	"net"
	"io"
	//"os"
)

// TCPServer struct
type TCPServer struct {
	Bind string
	Port int
}

// Start TCPServer
func main() {
	fmt.Printf("TCP Server\n")
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	defer ln.Close()
	if err != nil {
		panic(err)
	}
	


	for {
		conn, err := ln.Accept()
		defer conn.Close()
		fmt.Println("--- Accept")
		if err != nil {
			fmt.Println("---------1")
			panic(err)
		}

		buff := make([]byte, 1024)
		fmt.Println("---- Read")
		_, err = conn.Read(buff)
		fmt.Println("-------------2")
		if err != nil && err != io.EOF{
			fmt.Println(err)

			panic(err)
		} else {
			fmt.Println("recv:", string(buff))
			if err == io.EOF {
				fmt.Printf("This client EXIT\n")
				//os.Exit(0)
			}
		}
	}
}