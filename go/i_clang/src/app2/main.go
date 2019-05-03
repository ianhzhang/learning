package main


import (
	"fmt"

	"app2/myGolib/cl"
)

func main() {
	fmt.Println("Hello")
	platform := cl.Platform()
	fmt.Println("platform:", platform)
	fmt.Println("-------------------------------")
	device := cl.Device()
	fmt.Println("device:", device)


}