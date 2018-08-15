package pkg_send

import (
	"fmt" 
	"time"
)

func SenderTask(myChan chan int) {
	for i:=0; i< 10; i++ {
		fmt.Println("send:", i)
		myChan <- i
		time.Sleep(time.Millisecond * 2000)
	}
	fmt.Print("send loop end\n")
}
