package main

import (
	"fmt"
	"net"
)


func main() {
	fmt.Println("UDP Client")
    conn, err := net.Dial("udp", "127.0.0.1:8080")
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
	}
	defer conn.Close()
	
	msg := "Hello World!"
	buf := []byte(msg)
	_, err = conn.Write(buf)
	if err != nil {
		fmt.Println("error:", err)
	} 
}