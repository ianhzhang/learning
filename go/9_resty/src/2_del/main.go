package main

import (
	"fmt"
	"gopkg.in/resty.v1"
	"reflect"
)


func main() {
	fmt.Println("Go Resty")
	var url string = "http://localhost:3000/family/9"
	fmt.Println(url)
	resp, err := resty.R().Delete(url)
	fmt.Println(reflect.TypeOf(resp))
	fmt.Println(err)
	fmt.Println(resp)
}