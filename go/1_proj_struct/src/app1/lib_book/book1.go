package lib_book

// import other lib
import "app1/lib_greet"

import "fmt"

func Book1() {
	fmt.Println("Hello Book1 1")
	fmt.Println("** in book1")
	lib_greet.Greet1()
	fmt.Println("--- end book1")
}