package main

import (
	"fmt"
	"generic/container_api"
)



type A struct {
	A int
	B int
}

type B struct {
	A int 
	B int 
	C string
}

func main() {
	fmt.Println("Go Generic")

	list := container_api.List{}

	a1 := A{1,2}
	a2 := A{3,4}

	list.Append(a1)
	list = append(list, a2)
	fmt.Println(list)

	b1 := B{5,6,"Hello 1"}
	b2 := B{7,8,"Hello 3"}
	list.Append(b1)
	list = append(list, b2)
	fmt.Println(list)
	fmt.Println(list.GetByIndex(2))
}