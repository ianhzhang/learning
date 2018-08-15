package pkg_recv

import (
	"fmt" 
)

func RecvTask(myChan chan int) {
	for {
		msg := <- myChan
		fmt.Println("recv:",msg)
	}
	fmt.Print("recv loop end\n")
}