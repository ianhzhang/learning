package main

import (
	"fmt"
	"gopkg.in/resty.v1"
	"reflect"
	"encoding/json"
)


func main() {
	fmt.Println("Go Resty POST")
	name := map[string]interface{}{}
    //name["id"] = "8"
    name["fn"] = "Hong"
    name["ln"] = "Sun"

	nameStrByte,_ := json.Marshal(name)
	var url string = "http://localhost:3000/family"
	fmt.Println(url)

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		//SetBody(`{"id":"2","ln":"Sun","fn":"Hong"}`).
		SetBody(nameStrByte).
		Post(url)
		
	fmt.Println(reflect.TypeOf(resp))
	fmt.Println("\nError: %v", err)
	fmt.Println(resp)

}