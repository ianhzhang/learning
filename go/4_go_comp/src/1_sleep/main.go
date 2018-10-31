package main

import (
	"fmt"
	"time"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, 世界")
	logFile, err := os.OpenFile("/tmp/info.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err !=nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Print("Start")

	for i:=0; i<10; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i)
		log.Print(i)
	} 
}
