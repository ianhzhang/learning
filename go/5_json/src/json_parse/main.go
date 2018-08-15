package main

import (
	"fmt"
	"encoding/json"
	"reflect"
)

func main() {
	fmt.Println("Hello, 世界")
	//str := `{ "ln": "Zhang", "fn":"Ian" }`
	str := `{ "ln": "Zhang", "fn":"Ian","email":["izhang@extreme.com", "izhang@gmail.com"] }`

	var result map[string]interface{}
	
	_ = json.Unmarshal( []byte(str), &result);
	fmt.Println("result:", result)
	fmt.Println("ln:", result["ln"])
	fmt.Println("fn:", result["fn"])
	emails := result["email"]
	fmt.Println("emails:", emails)
	fmt.Println("emails type:", reflect.TypeOf(emails))
	result["new"]="new"
	
	//email0 := emails

	//fmt.Println((*email0)[0])

	fmt.Println("------------- End --------------")
}