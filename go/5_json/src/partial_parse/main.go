package main

import (
	"fmt"
	"encoding/json"
	"reflect"
)

type jsonStruct struct {
	Son struct {
		Ln string `json:fn`
		Fn string `json:ln`
	} `json:son`
}
func main() {
	fmt.Println("Hello world")
	//str := `{ "ln": "Zhang", "fn":"Ian" }`
	str := `{ "ln": "Zhang", "fn":"Ian","son":{"ln":"Zhang","fn":"Kevin"}, "email":["izhang@extreme.com", "izhang@gmail.com"] }`

	var jsonData jsonStruct
	_ = json.Unmarshal([]byte(str), &jsonData)
	fmt.Println(jsonData)
	fmt.Println(jsonData.Son)
	fmt.Println(reflect.TypeOf(jsonData))
	fmt.Println(reflect.TypeOf(jsonData.Son))
	fmt.Println(jsonData.Son.Ln)
	fmt.Println(jsonData.Son.Fn)

	fmt.Println("------------- End --------------")
}