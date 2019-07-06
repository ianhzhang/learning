package main

import(
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}
func main() {
	fmt.Println("Hello")
	t := T{23, "Ian"}

	s := reflect.ValueOf(&t).Elem()
	fmt.Println(s)

	for i:=0; i<s.NumField(); i++ {
		fmt.Println("i:", i,  ", type:", s.Field(i).Type(), ", field name:", s.Type().Field(i).Name, ", field value:", s.Field(i))
		/*
		i: 0 , type: int , field name: A , field value: 23
		i: 1 , type: string , field name: B , field value: Ian

		*/

	}

	attributes := s.Type()
	fmt.Println(attributes)
}