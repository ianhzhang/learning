package main

import (
	"fmt"
	"3_sys/restapi"
	"3_sys/hm"
)

func main() {
	fmt.Println("Start Main 1")
	go hm.Say("world")
	restapi.StartRest()
}