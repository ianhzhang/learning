package main

import (
	"fmt" 
	"time"
)

func f(n int) {
	for i:=0; i<n; i++ {
		fmt.Println(n,":", i)
		time.Sleep(time.Millisecond * 2000)
	}
	fmt.Print("loop end\n")
}

func main() {
	go f(10) 
	var input string
	fmt.Scanln(&input)
}