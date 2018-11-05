package main

import (
	"fmt"
	"encoding/json"
	//"reflect"
)

// Practice partial unmarshal
func main() {
	fmt.Println("Hello")

	// Full data
	type Data1 struct {
		A int `json "a"`
		B int `json "b"`
		C int `json "c"`
	}

	// Partial data
	type Data2 struct {
		A int `json "a"`
		B int `json "b"`
	}

	str := `{"a": 298202, "c": 258720, "b": 1234}`
	
	mydata1 := Data1{}
	_ = json.Unmarshal([]byte(str), &mydata1)

	fmt.Println(mydata1)

	mydata2 := Data2{}
	_ = json.Unmarshal([]byte(str), &mydata2)

	fmt.Println(mydata2)
	
}