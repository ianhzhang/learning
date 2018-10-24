package main

import (
	"fmt"
)

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (animal *Dog) Move(step int) {
	fmt.Printf("Dog %s moves %d steps\n", animal.name, step)
}


func (animal *Cat) Move(step int) {
	fmt.Printf("Cat %s moves %d steps\n", animal.name, step)
}

func hello() {
	fmt.Println("class hello")
}
