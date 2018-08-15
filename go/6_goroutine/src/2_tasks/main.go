package main

import (
	"fmt" 
	"time"
)

func receiver_task(myChan chan int) {
	for {
		msg := <- myChan
		fmt.Println("recv:",msg)
	}
	fmt.Print("recv loop end\n")
}

func sender_task(myChan chan int) {
	for i:=0; i< 10; i++ {
		fmt.Println("send:", i)
		myChan <- i
		time.Sleep(time.Millisecond * 2000)
	}
	fmt.Print("loop end\n")
}

func main() {
	var myChan chan int = make (chan int)
	go receiver_task(myChan)
	go sender_task(myChan)
	
	var input string
	fmt.Scanln(&input)
}