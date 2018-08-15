package main

import (
	"fmt"
	"2_tasks_1/pkg_send"
	"2_tasks_1/pkg_recv"
)


func main() {
	var myChan chan int = make (chan int)
	go pkg_recv.RecvTask(myChan)
	go pkg_send.SenderTask(myChan)
	
	var input string
	fmt.Scanln(&input)
}