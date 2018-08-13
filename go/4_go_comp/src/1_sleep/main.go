package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, 世界")
	
	for i:=0; i<10; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i)
	} 
}
