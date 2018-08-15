package lib_greet

import "fmt"

func Greet1() {
	fmt.Println("Hello Greet 1")
	// same package without import
	fmt.Println("------")
	fmt.Println("Same package without import")
	Greet2()
	fmt.Println("------")
}