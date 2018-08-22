package main

import (
	"fmt"
	"gopkg.in/resty.v1"
	"reflect"
)


func main() {
	fmt.Println("Go Resty")
	var url string = "http://localhost:3000/family"
	fmt.Println(url)
	resp, err := resty.R().Get(url)
	fmt.Println(reflect.TypeOf(resp))

	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Body: %v\n", resp)     // or resp.String() or string(resp.Body())

	url = "http://localhost:3000/family/1"
	fmt.Println(url)
	resp, err = resty.R().Get(url)
	fmt.Println(resp)
	body := resp.Body()
	fmt.Println(reflect.TypeOf(body))
	fmt.Println(body)
	// We can unmarshal body
}