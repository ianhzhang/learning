package main

import (
	"fmt"
	"net"
)

// TCPClient struct
type TCPClient struct {
	Host string
	Port int
}


func main() {
	fmt.Println("TCP Client")
    conn, err := net.Dial("tcp", "127.0.0.1:8080")
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
	}
	defer conn.Close()
	
	msg := "Hello World!"
	buf := []byte(msg)
	fmt.Println(buf)
	
	_, err = conn.Write(buf)
	if err != nil {
		fmt.Println("error:", err)
	}

	// if we do not wait, the tcp server will not receive.  It seems to me that it does not flush
	// But I tried bufio.Writer flush.  It still does not work.
	var input string
	fmt.Scanln(&input)
}