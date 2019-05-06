package main


import (
	"fmt"
	"reflect"
	"os"
	"io"
	"log"
)

func main() {
	logFile, _ := os.Create("ihzLogFile.txt")
	fmt.Println(reflect.TypeOf(logFile))
	fmt.Println("Hello")
	log.Println("Hello1")
	log.SetOutput(io.MultiWriter(os.Stderr, logFile))
	log.Println("Hello2")
	log.SetOutput(os.Stderr)
	log.Println("Hello3")
}